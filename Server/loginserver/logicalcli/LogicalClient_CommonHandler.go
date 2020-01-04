package logicalcli

import (
	"../../common/prot/pt_com"
	"../../common/prot/pt_logical"
	"../../common/prot/pt_login"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func (o *LogicalClient) SendRegisterNode() {
	data := pt_com.C2SRegisterNode{
		Nodetype: *proto.Uint32(uint32(pt_com.NodeType_NODE_LOGIN_SERVER)),
	}

	bytes, err := proto.Marshal(&data)
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

	o.channelStart <- data.Code == int32(pt_login.LoginCode_C_SUCCESS)
}

func (o *LogicalClient) SendUserLosted(loginsocket uint32) {
	data := pt_logical.S2LUserLost{
		Loginsocket: *proto.Uint32(loginsocket),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LogicalClient SendUserLost Error:", err.Error())
		return
	}

	o.SendPacket(uint32(pt_login.LoginPid_LOGIN_LOST), bytes)
}