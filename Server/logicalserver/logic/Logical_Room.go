package logic

import (
	"../../common"
	"../../common/prot/pt_com"
	"../../common/prot/pt_login"
	"../../common/tools"
	"../global"
	"../include"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func (o *Logical) LogicalNotifyRoomAppend(roomid uint64, roomname string, roomkind uint32) {
	component, ok := tools.GetComponentCollecter().GetComponent(common.LogicalServer)
	if !ok {
		return
	}
	logicalserver := component.(include.ILogicalServer)

	logins := *global.GetGlobalData().GetLogins()
	for k := range logins {
		logicalserver.NotifyRoomAppend(k, roomid, roomname, roomkind)
	}
}

func (o *Logical) LogicalNotifyRoomRemove(roomid uint64) {
	component, ok := tools.GetComponentCollecter().GetComponent(common.LogicalServer)
	if !ok {
		return
	}
	logicalserver := component.(include.ILogicalServer)

	logins := *global.GetGlobalData().GetLogins()
	for k := range logins {
		logicalserver.NotifyRoomRemoved(k, roomid)
	}
}

func (o *Logical) LogicalNotifyRoomUserAppend(roomid uint64, userid uint64, username string, gold uint64) {
	component, ok := tools.GetComponentCollecter().GetComponent(common.LogicalServer)
	if !ok {
		return
	}
	logicalserver := component.(include.ILogicalServer)

	r := global.GetGlobalData().GetRoomByRoomId(roomid)

	for k := range r.RoomUsers {
		userdata := global.GetGlobalData().GetUserData(k)
		logicalserver.NotifyRoomUserAppend(userdata.LogicalSocket, userdata.LoginSocket, userid, username, gold)
	}
}

func (o *Logical) LogicalNotifyRoomUserRemove(roomid uint64, userid uint64) {
	component, ok := tools.GetComponentCollecter().GetComponent(common.LogicalServer)
	if !ok {
		return
	}
	logicalserver := component.(include.ILogicalServer)

	r := global.GetGlobalData().GetRoomByRoomId(roomid)

	for k := range r.RoomUsers {
		userdata := global.GetGlobalData().GetUserData(k)
		logicalserver.NotifyRoomUserRemove(userdata.LogicalSocket, userdata.LoginSocket, userid)
	}
}

func (o *Logical) LogicalUserEnterRoom(params ...interface{}) {
	fmt.Println(params...)

	logicalsocket := params[0].(uint32)
	loginsocket := params[1].(uint32)
	roomid := params[2].(uint64)

	component, ok := tools.GetComponentCollecter().GetComponent(common.LogicalServer)
	if !ok {
		return
	}
	logicalserver := component.(include.ILogicalServer)

	if !global.GetGlobalData().IsRoomExist(roomid) {
		logicalserver.SendUserEnterRoomResult(logicalsocket, loginsocket, int32(pt_login.LoginCode_C_ROOM_NOT_EXIST), roomid, nil, nil)
		return
	}

	if global.GetGlobalData().IsRoomUserBySocket(logicalsocket, loginsocket) {
		logicalserver.SendUserEnterRoomResult(logicalsocket, loginsocket, int32(pt_login.LoginCode_C_ALREADY_IN_ROOM), roomid, nil, nil)
		return
	}

	userid := global.GetGlobalData().GetUserIdBySocket(logicalsocket, loginsocket)
	global.GetGlobalData().AddRoomUser(roomid, userid)
	roomdata := global.GetGlobalData().GetRoomByRoomId(roomid)

	users := make(map[uint64]*pt_com.RoomUserItem)
	for k, v := range roomdata.RoomUsers {
		userdata := global.GetGlobalData().GetUserData(k)
		users[k] = &pt_com.RoomUserItem{
			Userid: *proto.Uint64(userdata.UserID),
			Name: *proto.String(userdata.UserName),
			Gold: *proto.Uint64(userdata.Gold),
			Tableid: *proto.Uint32(v.TableID),
			Chairid: *proto.Uint32(v.ChairID),
		}
	}

	tables := make(map[uint32]*pt_com.RoomTableItem)
	for k, v := range roomdata.RoomTables {
		tables[k] = &pt_com.RoomTableItem{
			Tableid: *proto.Uint32(v.TableID),
			Capacity: *proto.Uint32(v.Capacity),
		}
	}

	logicalserver.SendUserEnterRoomResult(logicalsocket, loginsocket, int32(pt_login.LoginCode_C_SUCCESS), roomid, users, tables)

	userdata := global.GetGlobalData().GetUserData(userid)
	o.LogicalNotifyRoomUserAppend(roomid, userdata.UserID, userdata.UserName, userdata.Gold)
}

func (o *Logical) LogicalUserLeaveRoom(params ...interface{}) {
	fmt.Println(params...)

	component, ok := tools.GetComponentCollecter().GetComponent(common.LogicalServer)
	if !ok {
		return
	}
	logicalserver := component.(include.ILogicalServer)

	logicalsocket := params[0].(uint32)
	loginsocket := params[1].(uint32)

	if !global.GetGlobalData().IsRoomUserBySocket(logicalsocket, loginsocket) {
		logicalserver.SendUserLeaveRoomResult(logicalsocket, loginsocket, int32(pt_login.LoginCode_C_NOT_IN_ROOM), 0)
		return
	}

	userid := global.GetGlobalData().GetUserIdBySocket(logicalsocket, loginsocket)
	roomid := global.GetGlobalData().RemoveRoomUser(userid)
	logicalserver.SendUserLeaveRoomResult(logicalsocket, loginsocket, int32(pt_login.LoginCode_C_NOT_IN_ROOM), roomid)

	o.LogicalNotifyRoomUserRemove(roomid, userid)
}