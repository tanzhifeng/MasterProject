package logicalcli

import (
	"../../common"
	"../../common/prot/pt_logical"
	"../../common/prot/pt_login"
	"../logic"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func (o *LogicalClient) onNotifyRoomAppend(bytes []byte) {
	var data pt_login.S2CNotifyAppendRoom
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LogicalClient onNotifyRoomAppend Error: ", err.Error())
		return
	}

	fmt.Println("LogicalClient onNotifyRoomAppend :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_NOTIFY_APPEND_ROOM, data.Item.Roomid, data.Item.Roomname, data.Item.Roomkind)
}

func (o *LogicalClient) onNotifyRoomRemove(bytes []byte) {
	var data pt_login.S2CNotifyRemoveRoom
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LogicalClient onNotifyRoomRemove Error: ", err.Error())
		return
	}

	fmt.Println("LogicalClient onNotifyRoomRemove :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_NOTIFY_REMOVE_ROOM, data.Roomid)
}

func (o *LogicalClient) SendUserEnterRoom(loginsocket uint32, roomid uint64) {
	data := pt_logical.S2LRoomUserEnter{
		Loginsocket: *proto.Uint32(loginsocket),
		Roomid: *proto.Uint64(roomid),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LogicalClient SendUserEnterRoom Error:", err.Error())
		return
	}

	o.SendPacket(uint32(pt_login.LoginPid_LOGIN_ROOM_USER_ENTER), bytes)
}

func (o *LogicalClient) SendUserLeaveRoom(loginsocket uint32) {
	data := pt_logical.S2LRoomUserLeave{
		Loginsocket: *proto.Uint32(loginsocket),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LogicalClient SendUserLeaveRoom Error:", err.Error())
		return
	}

	o.SendPacket(uint32(pt_login.LoginPid_LOGIN_ROOM_USER_LEAVE), bytes)
}

func (o *LogicalClient) onUserEnterRoomResult(bytes []byte) {
	var data pt_logical.L2SRoomUserEnterResult
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LogicalClient onUserEnterRoomResult Error: ", err.Error())
		return
	}

	fmt.Println("LogicalClient onUserEnterRoomResult :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_USER_ENTER_ROOM_RESULT, data.Loginsocket, data.Code, data.Roomid, data.Users, data.Tables)
}

func (o *LogicalClient) onUserLeaveRoomResult(bytes []byte) {
	var data pt_logical.L2SRoomUserLeaveResult
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LogicalClient onUserLeaveRoomResult Error: ", err.Error())
		return
	}

	fmt.Println("LogicalClient onUserLeaveRoomResult :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_USER_LEAVE_ROOM_RESULT, data.Loginsocket, data.Code, data.Roomid)
}

func (o *LogicalClient) onNotifyRoomUserAppend(bytes []byte) {
	var data pt_logical.L2SRoomNotifyAppendUser
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LogicalClient onNotifyRoomUserAppend Error: ", err.Error())
		return
	}

	fmt.Println("LogicalClient onNotifyRoomUserAppend :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_NOTIFY_APPEND_ROOM_USER, data.Loginsocket, data.Item.Userid, data.Item.Name, data.Item.Gold, data.Item.Tableid, data.Item.Chairid)
}

func (o *LogicalClient) onNotifyRoomUserRemove(bytes []byte) {
	var data pt_logical.L2SRoomNotifyRemoveUser
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LogicalClient onNotifyRoomUserRemove Error: ", err.Error())
		return
	}

	fmt.Println("LogicalClient onNotifyRoomUserRemove :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_NOTIFY_REMOVE_ROOM_USER, data.Loginsocket, data.Userid)
}