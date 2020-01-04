package logicalsvr

import (
	"../../common"
	"../../common/prot/pt_com"
	"../../common/prot/pt_logical"
	"../../common/prot/pt_login"
	"../logic"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func (o *LogicalServer) NotifyRoomAppend(logicalsocket uint32, roomid uint64, roomname string, roomkind uint32) {
	data := pt_login.S2CNotifyAppendRoom{
		Item: &pt_com.RoomItem{
			Roomid: *proto.Uint64(roomid),
			Roomname: *proto.String(roomname),
			Roomkind: *proto.Uint32(roomkind),
		},
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LogicalServer NotifyRoomAppend Error:", err.Error())
		return
	}

	o.SendPacket(logicalsocket, uint32(pt_login.LoginPid_LOGIN_ROOM_NOTIFY_APPEND_ROOM), bytes)
}

func (o *LogicalServer) NotifyRoomRemoved(logicalsocket uint32, roomid uint64) {
	data := pt_login.S2CNotifyRemoveRoom{
		Roomid: *proto.Uint64(roomid),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LogicalServer NotifyRoomRemoved Error:", err.Error())
		return
	}

	o.SendPacket(logicalsocket, uint32(pt_login.LoginPid_LOGIN_ROOM_NOTIFY_REMOVE_ROOM), bytes)
}

func (o *LogicalServer) NotifyRoomUserAppend(logicalsocket uint32, loginsocket uint32, userid uint64, username string, gold uint64) {
	data := pt_logical.L2SRoomNotifyAppendUser{
		Loginsocket: *proto.Uint32(loginsocket),
		Item: &pt_com.RoomUserItem{
			Userid: *proto.Uint64(userid),
			Name: *proto.String(username),
			Gold: *proto.Uint64(gold),
		},
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LogicalServer NotifyRoomUserAppend Error:", err.Error())
		return
	}

	o.SendPacket(logicalsocket, uint32(pt_login.LoginPid_LOGIN_ROOM_NOTIFY_APPEND_USER), bytes)
}

func (o *LogicalServer) NotifyRoomUserRemove(logicalsocket uint32, loginsocket uint32, userid uint64) {
	data := pt_logical.L2SRoomNotifyRemoveUser{
		Loginsocket: *proto.Uint32(loginsocket),
		Userid: *proto.Uint64(userid),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LogicalServer NotifyRoomUserLeave Error:", err.Error())
		return
	}

	o.SendPacket(logicalsocket, uint32(pt_login.LoginPid_LOGIN_ROOM_NOTIFY_REMOVE_USER), bytes)
}

func (o *LogicalServer) SendUserEnterRoomResult(logicalsocket uint32, loginsocket uint32, code int32, roomid uint64, users map[uint64]*pt_com.RoomUserItem, tables map[uint32]*pt_com.RoomTableItem) {
	data := pt_logical.L2SRoomUserEnterResult{
		Loginsocket: *proto.Uint32(loginsocket),
		Code: *proto.Int32(code),
		Roomid: *proto.Uint64(roomid),
		Users: users,
		Tables: tables,
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LogicalServer SendUserEnterRoomResult Error:", err.Error())
		return
	}

	o.SendPacket(logicalsocket, uint32(pt_login.LoginPid_LOGIN_ROOM_USER_ENTER), bytes)
}

func (o *LogicalServer) SendUserLeaveRoomResult(logicalsocket uint32, loginsocket uint32, code int32, roomid uint64) {
	data := pt_logical.L2SRoomUserLeaveResult{
		Loginsocket: *proto.Uint32(loginsocket),
		Code: *proto.Int32(code),
		Roomid: *proto.Uint64(roomid),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LogicalServer SendUserLeaveRoomResult Error:", err.Error())
		return
	}

	o.SendPacket(logicalsocket, uint32(pt_login.LoginPid_LOGIN_ROOM_USER_LEAVE), bytes)
}

func (o *LogicalServer) onUserEnterRoom(socket uint32, bytes []byte) {
	var data pt_logical.S2LRoomUserEnter
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LogicalServer onUserEnterRoom Error: ", err.Error())
		return
	}

	fmt.Println("LogicalServer onUserEnterRoom :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_USER_ENTER_ROOM, socket, data.Loginsocket, data.Roomid)
}

func (o *LogicalServer) onUserLeaveRoom(socket uint32, bytes []byte) {
	var data pt_logical.S2LRoomUserLeave
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LogicalServer onUserLeaveRoom Error: ", err.Error())
		return
	}

	fmt.Println("LogicalServer onUserLeaveRoom :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_USER_LEAVE_ROOM, socket, data.Loginsocket)
}