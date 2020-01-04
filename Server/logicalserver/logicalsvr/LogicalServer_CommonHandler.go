package logicalsvr

import (
	"../../common"
	"../../common/prot/pt_com"
	"../../common/prot/pt_logical"
	"../logic"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
)

func (o *LogicalServer) SendRegisterNodeResult(logicalsocket uint32, code int32, nodetype uint32, extra *any.Any) {
	data := pt_com.S2CRegisterNodeResult{
		Code: *proto.Int32(code),
		Nodetype: *proto.Uint32(nodetype),
		Extra: extra,
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LogicalServer SendRegisterNodeResult Error:", err.Error())
		return
	}

	o.SendPacket(logicalsocket, uint32(pt_com.ComPid_COM_REGISTER_NODE), bytes)
}

func (o *LogicalServer) onRegisterNode(socket uint32, bytes []byte) {
	var data pt_com.C2SRegisterNode
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LogicalServer onRegisterNode Error: ", err.Error())
		return
	}

	fmt.Println("LogicalServer onRegisterNode :", data)

	logic.GetLogical().AppendEvent(common.EVENT_COMMON_REGISTER_NODE, socket, data.Nodetype, data.Extra)
}

func (o *LogicalServer) onUserLosted(socket uint32, bytes []byte) {
	var data pt_logical.S2LUserLost
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LogicalServer onUserLosted Error: ", err.Error())
		return
	}

	fmt.Println("LogicalServer onUserLosted :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_LOSTED, socket, data.Loginsocket)
}