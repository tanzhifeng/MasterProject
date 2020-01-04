package logicalcli

import (
	"../../common"
	"../../common/prot/pt_logical"
	"../../common/prot/pt_login"
	"../logic"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func (o *LogicalClient) SendUserLogin(loginsocket uint32, address string, account string, password string) {
	data := pt_logical.S2LLoginUser{
		Loginsocket: *proto.Uint32(loginsocket),
		Address: *proto.String(address),
		Account: *proto.String(account),
		Password: *proto.String(password),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LogicalClient SendUserLogin Error:", err.Error())
		return
	}

	o.SendPacket(uint32(pt_login.LoginPid_LOGIN_LOGIN), bytes)
}

func (o *LogicalClient) SendUserRegister(loginsocket uint32, account string, password string) {
	data := pt_logical.S2LRegisterUser {
		Loginsocket: *proto.Uint32(loginsocket),
		Account: *proto.String(account),
		Password: *proto.String(password),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LogicalClient SendUserRegister Error:", err.Error())
		return
	}

	o.SendPacket(uint32(pt_login.LoginPid_LOGIN_REGISTER), bytes)
}

func (o *LogicalClient) onUserLoginResult(bytes []byte) {
	var data pt_logical.L2SLoginUserResult
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LogicalClient onUserLoginResult Error: ", err.Error())
		return
	}

	fmt.Println("LogicalClient onUserLoginResult :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_LOGIN_RESULT, data.Loginsocket, data.Code, data.Userid, data.Account, data.Password, data.Gold, data.Rooms)
}

func (o *LogicalClient) onUserRegisterResult(bytes []byte) {
	var data pt_logical.L2SRegisterUserResult
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LogicalClient onUserRegisterResult Error: ", err.Error())
		return
	}

	fmt.Println("LogicalClient onUserRegisterResult :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_REGISTER_RESULT, data.Loginsocket, data.Code, data.Account, data.Password)
}