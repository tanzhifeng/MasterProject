package logic

import (
	"../../common"
	"../../common/prot/pt_com"
	"../../common/prot/pt_login"
	"../../common/tools"
	"../global"
	"../include"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

func (o *Logical) LogicalRegisterNode(params ...interface{}) {
	fmt.Println(params...)

	component, ok := tools.GetComponentCollecter().GetComponent(common.LogicalServer)
	if !ok {
		return
	}

	logicalsocket := params[0].(uint32)
	nodetype := params[1].(uint32)
	extra := params[2].(*any.Any)

	switch nodetype {
	case uint32(pt_com.NodeType_NODE_LOGIN_SERVER):
		global.GetGlobalData().AddLoginServerNode(logicalsocket)
	case uint32(pt_com.NodeType_NODE_ROOM_SERVER):
		var roomdata pt_com.RoomItem
		err := ptypes.UnmarshalAny(extra, &roomdata)
		if err != nil {
			fmt.Println("Logical LogicalRegisterNode Error: ", err.Error())
			return
		}
		global.GetGlobalData().AddRoomServerNode(logicalsocket, roomdata.Roomid, roomdata.Roomname, roomdata.Roomkind)
		o.LogicalNotifyRoomAppend(roomdata.Roomid, roomdata.Roomname, roomdata.Roomkind)
	default:
		fmt.Printf("Unknow NodeType [%d]\n", nodetype)
		return
	}

	logicalserver := component.(include.ILogicalServer)
	logicalserver.SendRegisterNodeResult(logicalsocket, int32(pt_login.LoginCode_C_SUCCESS), nodetype, extra)
}

func (o *Logical) LogicalServerConnected(params ...interface{}) {
	fmt.Println(params...)

	logicalsocket := params[0].(uint32)
	address := params[1].(string)

	global.GetGlobalData().AddNodeClient(logicalsocket, address)
}
//登录服掉线
func (o *Logical) LogicalLoginServerLosted(logicalsocket uint32) {
	sockets := *global.GetGlobalData().GetLoginSockets(logicalsocket)

	for _, userid := range sockets {
		v := global.GetGlobalData().GetUserData(userid)
		o.LogicalUserLosted(v.LogicalSocket, v.LoginSocket)
	}
	global.GetGlobalData().RemoveLoginServerNode(logicalsocket)
}
//房间服掉线
func (o *Logical) LogicalRoomServerLosted(logicalsocket uint32) {
	roomdata := global.GetGlobalData().GetRoomBySocket(logicalsocket)

	global.GetGlobalData().RemoveRoomServerNode(logicalsocket)
	o.LogicalNotifyRoomRemove(roomdata.RoomID)
}

func (o *Logical) LogicalServerLosted(params ...interface{}) {
	fmt.Println(params...)

	logicalsocket := params[0].(uint32)
	clientdata := global.GetGlobalData().GetNodeClient(logicalsocket)

	switch clientdata.NodeType {
	case uint32(pt_com.NodeType_NODE_LOGIN_SERVER):
		o.LogicalLoginServerLosted(logicalsocket)
	case uint32(pt_com.NodeType_NODE_ROOM_SERVER):
		o.LogicalRoomServerLosted(logicalsocket)
	case uint32(pt_com.NodeType_NODE_UNKNOW):
		fmt.Println("UnknowServer Losted !")
	}
}

func (o *Logical) LogicalUserLosted(params ...interface{}) {
	fmt.Println(params...)

	logicalsocket := params[0].(uint32)
	loginsocket := params[1].(uint32)

	if global.GetGlobalData().IsUserOnlineBySocket(logicalsocket, loginsocket) {
		userid := global.GetGlobalData().GetUserIdBySocket(logicalsocket, loginsocket)
		v := global.GetGlobalData().GetUserData(userid)

		component, ok := tools.GetComponentCollecter().GetComponent(common.DatabaseClient)
		if !ok {
			return
		}

		databasecli := component.(include.IDatabaseClient)

		switch v.Status {
		case common.USER_FREE://玩家空闲,直接保存玩家数据
			databasecli.SendUserDataSave(v.UserID, v.Gold)
			global.GetGlobalData().RemoveUserByUserId(userid)
		case common.USER_GAME://玩家游戏中,进入托管状态

		case common.USER_TRUST://玩家托管中,继续托管

		}
	}
}