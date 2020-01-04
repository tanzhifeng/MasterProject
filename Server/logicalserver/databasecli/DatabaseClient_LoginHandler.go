package databasecli

import (
	"../../common"
	"../../common/prot/pt_database"
	"../../common/prot/pt_login"
	"../logic"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func (o *DatabaseClient) SendUserLogin(logicalsocket uint32, loginsocket uint32, address string, account string, password string) {
	data := pt_database.L2DLoginUser{
		Logicalsocket: *proto.Uint32(logicalsocket),
		Loginsocket: *proto.Uint32(loginsocket),
		Address: *proto.String(address),
		Account: *proto.String(account),
		Password: *proto.String(password),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("DatabaseClient SendUserLogin Error:", err.Error())
		return
	}

	o.SendPacket(uint32(pt_login.LoginPid_LOGIN_LOGIN), bytes)
}

func (o *DatabaseClient) SendUserRegister(logicalsocket uint32, loginsocket uint32, account string, password string) {
	data := pt_database.L2DRegisterUser {
		Logicalsocket: *proto.Uint32(logicalsocket),
		Loginsocket: *proto.Uint32(loginsocket),
		Account: *proto.String(account),
		Password: *proto.String(password),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("DatabaseClient SendUserRegister Error:", err.Error())
		return
	}

	o.SendPacket(uint32(pt_login.LoginPid_LOGIN_REGISTER), bytes)
}

func (o *DatabaseClient) onUserLoginResult(bytes []byte) {
	var data pt_database.D2LLoginUserResult
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("DatabaseClient onUserLoginResult Error: ", err.Error())
		return
	}

	fmt.Println("DatabaseClient onUserLoginResult :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_LOGIN_RESULT, data.Logicalsocket, data.Loginsocket, data.Address, data.Code, data.Userid, data.Account, data.Password, data.Gold)
}

func (o *DatabaseClient) onUserRegisterResult(bytes []byte) {
	var data pt_database.D2LRegisterUserResult
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("DatabaseClient onUserRegisterResult Error: ", err.Error())
		return
	}

	fmt.Println("DatabaseClient onUserRegisterResult :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_REGISTER_RESULT, data.Logicalsocket, data.Loginsocket, data.Code, data.Account, data.Password)
}