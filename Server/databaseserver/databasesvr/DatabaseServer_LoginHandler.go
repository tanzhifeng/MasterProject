package databasesvr

import (
	"../../common"
	"../../common/prot/pt_database"
	"../../common/prot/pt_login"
	"../logic"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func (o *DatabaseServer) SendUserLoginResult(socket uint32, code int32, logicalsocket uint32, loginsocket uint32, address string, userid uint64, account string, password string, gold uint64) {
	data := pt_database.D2LLoginUserResult{
		Logicalsocket: *proto.Uint32(logicalsocket),
		Loginsocket: *proto.Uint32(loginsocket),
		Address: *proto.String(address),
		Code: *proto.Int32(code),
		Userid: *proto.Uint64(userid),
		Account: *proto.String(account),
		Password: *proto.String(password),
		Gold: *proto.Uint64(gold),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("DatabaseServer SendUserLoginResult Error:", err.Error())
		return
	}

	o.SendPacket(socket, uint32(pt_login.LoginPid_LOGIN_LOGIN), bytes)
}

func (o *DatabaseServer) SendUserRegisterResult(socket uint32, code int32, logicalsocket uint32, loginsocket uint32, account string, password string) {
	data := pt_database.D2LRegisterUserResult{
		Logicalsocket: *proto.Uint32(logicalsocket),
		Loginsocket: *proto.Uint32(loginsocket),
		Code: *proto.Int32(code),
		Account: *proto.String(account),
		Password: *proto.String(password),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("DatabaseServer SendUserRegisterResult Error:", err.Error())
		return
	}

	o.SendPacket(socket, uint32(pt_login.LoginPid_LOGIN_REGISTER), bytes)
}

func (o *DatabaseServer) onUserLogin(socket uint32, bytes []byte) {
	var data pt_database.L2DLoginUser
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("DatabaseServer onUserLogin Error: ", err.Error())
		return
	}

	fmt.Println("DatabaseServer onUserLogin :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_LOGIN, socket, data.Logicalsocket, data.Loginsocket, data.Address, data.Account, data.Password)
}

func (o *DatabaseServer) onUserRegister(socket uint32, bytes []byte) {
	var data pt_database.L2DRegisterUser
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("DatabaseServer onUserRegister Error: ", err.Error())
		return
	}

	fmt.Println("DatabaseServer onUserRegister :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_REGISTER, socket, data.Logicalsocket, data.Loginsocket, data.Account, data.Password)
}