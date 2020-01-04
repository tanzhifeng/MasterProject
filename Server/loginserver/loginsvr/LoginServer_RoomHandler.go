package loginsvr

import (
	"fmt"
	"github.com/golang/protobuf/proto"

	"../../common"
	"../../common/prot/pt_com"
	"../../common/prot/pt_login"
	"../logic"
)

func (o *LoginServer) NotifyRoomAppend(loginsocket uint32, roomid uint64, roomname string, roomkind uint32) {
	data := pt_login.S2CNotifyAppendRoom{
		Item: &pt_com.RoomItem{
			Roomid: *proto.Uint64(roomid),
			Roomname: *proto.String(roomname),
			Roomkind: *proto.Uint32(roomkind),
		},
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LoginServer NotifyRoomAppend Error:", err.Error())
		return
	}

	o.SendPacket(loginsocket, uint32(pt_login.LoginPid_LOGIN_ROOM_NOTIFY_APPEND_ROOM), bytes)
}

func (o *LoginServer) NotifyRoomRemove(loginsocket uint32, roomid uint64) {
	data := pt_login.S2CNotifyRemoveRoom{
		Roomid: *proto.Uint64(roomid),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LoginServer NotifyRoomRemove Error:", err.Error())
		return
	}

	o.SendPacket(loginsocket, uint32(pt_login.LoginPid_LOGIN_ROOM_NOTIFY_REMOVE_ROOM), bytes)
}

func (o *LoginServer) NotifyRoomUserAppend(loginsocket uint32, userid uint64, username string, gold uint64, tableid uint32, chairid uint32) {
	data := pt_login.S2CRoomNotifyAppendUser{
		Item: &pt_com.RoomUserItem{
			Userid: *proto.Uint64(userid),
			Name: *proto.String(username),
			Gold: *proto.Uint64(gold),
			Tableid: *proto.Uint32(tableid),
			Chairid: *proto.Uint32(chairid),
		},
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LoginServer NotifyRoomUserAppend Error:", err.Error())
		return
	}

	o.SendPacket(loginsocket, uint32(pt_login.LoginPid_LOGIN_ROOM_NOTIFY_APPEND_USER), bytes)
}

func (o *LoginServer) NotifyRoomUserRemove(loginsocket uint32, userid uint64) {
	data := pt_login.S2CRoomNotifyRemoveUser{
		Userid: *proto.Uint64(userid),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LoginServer NotifyRoomUserRemove Error:", err.Error())
		return
	}

	o.SendPacket(loginsocket, uint32(pt_login.LoginPid_LOGIN_ROOM_NOTIFY_REMOVE_USER), bytes)
}

func (o *LoginServer) SendUserEnterRoomResult(loginsocket uint32, code int32, roomid uint64, users map[uint64]*pt_com.RoomUserItem, tables map[uint32]*pt_com.RoomTableItem) {
	data := pt_login.S2CRoomUserEnterResult{
		Code: *proto.Int32(code),
		Roomid: *proto.Uint64(roomid),
		Users: users,
		Tables: tables,
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LoginServer SendUserEnterRoomResult Error:", err.Error())
		return
	}

	o.SendPacket(loginsocket, uint32(pt_login.LoginPid_LOGIN_ROOM_USER_ENTER), bytes)
}

func (o *LoginServer) SendUserLeaveRoomResult(loginsocket uint32, code int32, roomid uint64) {
	data := pt_login.S2CRoomUserLeaveResult{
		Code: *proto.Int32(code),
		Roomid: *proto.Uint64(roomid),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LoginServer SendUserLeaveRoomResult Error:", err.Error())
		return
	}

	o.SendPacket(loginsocket, uint32(pt_login.LoginPid_LOGIN_ROOM_USER_LEAVE), bytes)
}

func (o *LoginServer) onUserEnterRoom(socket uint32, bytes []byte) {
	var data pt_login.C2SRoomUserEnter
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LoginServer onUserEnterRoom Error: ", err.Error())
		return
	}

	fmt.Println("LoginServer onUserEnterRoom :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_USER_ENTER_ROOM, socket, data.Roomid)
}

func (o *LoginServer) onUserLeaveRoom(socket uint32, bytes []byte) {
	fmt.Println("LoginServer onUserLeaveRoom !")

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_USER_LEAVE_ROOM, socket)
}