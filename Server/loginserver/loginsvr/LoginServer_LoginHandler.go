package loginsvr

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"time"

	"../../common"
	"../../common/prot/pt_com"
	"../../common/prot/pt_login"
	"../logic"
)

func (o *LoginServer) SendUserLoginResult(loginsocket uint32, code int32, userid uint64, account string, password string, gold uint64, rooms map[uint64]*pt_com.RoomItem) {
	data := pt_login.S2CLoginUserResult{
		Code: *proto.Int32(code),
		Time: *proto.Int64(time.Now().Unix()),
		Userid: *proto.Uint64(userid),
		Account: *proto.String(account),
		Password: *proto.String(password),
		Gold: *proto.Uint64(gold),
		Rooms: rooms,
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LoginServer SendUserLoginResult Error:", err.Error())
		return
	}

	o.SendPacket(loginsocket, uint32(pt_login.LoginPid_LOGIN_LOGIN), bytes)
}

func (o *LoginServer) SendUserRegisterResult(loginsocket uint32, code int32, account string, password string) {
	data := pt_login.S2CRegisterUserResult{
		Code: *proto.Int32(code),
		Account: *proto.String(account),
		Password: *proto.String(password),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LoginServer SendLoginClientRegisterResult Error:", err.Error())
		return
	}

	o.SendPacket(loginsocket, uint32(pt_login.LoginPid_LOGIN_REGISTER), bytes)
}

func (o *LoginServer) onUserLogin(socket uint32, bytes []byte) {
	var data pt_login.C2SLoginUser
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LoginServer onUserLogin Error: ", err.Error())
		return
	}

	fmt.Println("LoginServer onUserLogin :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_LOGIN, socket, data.Account, data.Password)
}

func (o *LoginServer) onUserRegister(socket uint32, bytes []byte) {
	var data pt_login.C2SRegisterUser
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LoginServer onUserRegister Error: ", err.Error())
		return
	}

	fmt.Println("LoginServer onUserRegister :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_REGISTER, socket, data.Account, data.Password)
}