package logicalcli

import (
	"../../common/prot/pt_com"
	"../../common/prot/pt_login"
	"../logic"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

func (o *LogicalClient) SendRegisterNode() {
	config := *logic.GetLogical().GetConfig()

	roomid := uint64(config["roomid"].(float64))
	roomname := config["roomname"].(string)
	roomkind := uint32(config["roomkind"].(float64))

	roomdata := pt_com.RoomItem{
		Roomid: *proto.Uint64(roomid),
		Roomname: *proto.String(roomname),
		Roomkind: *proto.Uint32(roomkind),
	}
	bytes, err := proto.Marshal(&roomdata)
	if err != nil {
		fmt.Println("LogicalClient SendRegisterNode Error:", err.Error())
		return
	}

	data := pt_com.C2SRegisterNode{
		Nodetype: *proto.Uint32(uint32(pt_com.NodeType_NODE_ROOM_SERVER)),
		Extra: &any.Any{
			TypeUrl: "./pt_com.room_item",
			Value: bytes,
		},
	}

	bytes, err = proto.Marshal(&data)
	if err != nil {
		fmt.Println("LogicalClient SendRegisterNode Error:", err.Error())
		return
	}

	o.SendPacket(uint32(pt_com.ComPid_COM_REGISTER_NODE), bytes)
}

func (o *LogicalClient) onRegisterNodeResult(bytes []byte) {
	var data pt_com.S2CRegisterNodeResult
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LogicalClient onRegisterNodeResult Error: ", err.Error())
		return
	}

	var roomdata pt_com.RoomItem
	err = ptypes.UnmarshalAny(data.Extra, &roomdata)
	if err != nil {
		fmt.Println("Logical LogicalRegisterNode Error: ", err.Error())
		return
	}

	o.channelStart <- data.Code == int32(pt_login.LoginCode_C_SUCCESS)
}