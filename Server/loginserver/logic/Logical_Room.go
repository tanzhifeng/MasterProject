package logic

import (
	"../../common"
	"../../common/prot/pt_com"
	"../../common/tools"
	"../global"
	"../include"
	"fmt"
)

func (o *Logical) LoginServerNotifyAppendRoom(params ...interface{}) {
	fmt.Println(params...)

	component, ok := tools.GetComponentCollecter().GetComponent(common.LoginServer)
	if !ok {
		return
	}
	loginserver := component.(include.ILoginServer)

	roomid := params[0].(uint64)
	roomname := params[1].(string)
	roomkind := params[2].(uint32)

	accounts := *global.GetGlobalData().GetAccounts()

	for _, socket := range accounts {
		loginserver.NotifyRoomAppend(socket, roomid, roomname, roomkind)
	}
}

func (o *Logical) LoginServerNotifyRemoveRoom(params ...interface{}) {
	fmt.Println(params...)

	component, ok := tools.GetComponentCollecter().GetComponent(common.LoginServer)
	if !ok {
		return
	}
	loginserver := component.(include.ILoginServer)

	roomid := params[0].(uint64)

	accounts := *global.GetGlobalData().GetAccounts()

	for _, socket := range accounts {
		loginserver.NotifyRoomRemove(socket, roomid)
	}
}

func (o *Logical) LoginServerNotifyAppendRoomUser(params ...interface{}) {
	fmt.Println(params...)

	component, ok := tools.GetComponentCollecter().GetComponent(common.LoginServer)
	if !ok {
		return
	}
	loginserver := component.(include.ILoginServer)

	loginsocket := params[0].(uint32)
	userid := params[1].(uint64)
	username := params[2].(string)
	gold := params[3].(uint64)
	tableid := params[4].(uint32)
	chairid := params[5].(uint32)

	loginserver.NotifyRoomUserAppend(loginsocket, userid, username, gold, tableid, chairid)
}

func (o *Logical) LoginServerNotifyRemoveRoomUser(params ...interface{}) {
	fmt.Println(params...)

	component, ok := tools.GetComponentCollecter().GetComponent(common.LoginServer)
	if !ok {
		return
	}
	loginserver := component.(include.ILoginServer)

	loginsocket := params[0].(uint32)
	userid := params[1].(uint64)

	loginserver.NotifyRoomUserRemove(loginsocket, userid)
}

func (o *Logical) LoginServerUserEnterRoom(params ...interface{}) {
	fmt.Println(params...)

	loginsocket := params[0].(uint32)
	roomid := params[1].(uint64)

	component, ok := tools.GetComponentCollecter().GetComponent(common.LogicalClient)
	if !ok {
		return
	}
	logicalcli := component.(include.ILogicalClient)
	logicalcli.SendUserEnterRoom(loginsocket, roomid)
}

func (o *Logical) LoginServerUserEnterRoomResult(params ...interface{}) {
	fmt.Println(params...)

	loginsocket := params[0].(uint32)
	code := params[1].(int32)
	roomid := params[2].(uint64)
	users := params[3].(map[uint64]*pt_com.RoomUserItem)
	tables := params[4].(map[uint32]*pt_com.RoomTableItem)

	component, ok := tools.GetComponentCollecter().GetComponent(common.LoginServer)
	if !ok {
		return
	}
	loginserver := component.(include.ILoginServer)
	loginserver.SendUserEnterRoomResult(loginsocket, code, roomid, users, tables)
}

func (o *Logical) LoginServerUserLeaveRoom(params ...interface{}) {
	fmt.Println(params...)

	loginsocket := params[0].(uint32)

	component, ok := tools.GetComponentCollecter().GetComponent(common.LogicalClient)
	if !ok {
		return
	}
	logicalcli := component.(include.ILogicalClient)
	logicalcli.SendUserLeaveRoom(loginsocket)
}

func (o *Logical) LoginServerUserLeaveRoomResult(params ...interface{}) {
	fmt.Println(params...)

	loginsocket := params[0].(uint32)
	code := params[1].(int32)
	roomid := params[2].(uint64)

	component, ok := tools.GetComponentCollecter().GetComponent(common.LoginServer)
	if !ok {
		return
	}
	loginserver := component.(include.ILoginServer)
	loginserver.SendUserLeaveRoomResult(loginsocket, code, roomid)
}