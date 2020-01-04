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

func (o *LogicalServer) SendUserLoginResult(logicalsocket uint32, loginsocket uint32, address string, retcode int32, userid uint64, account string, password string, gold uint64, rooms *map[uint64]*common.DataRoom) {
	rm := make(map[uint64]*pt_com.RoomItem)
	for k, v := range *rooms {
		rm[k] = &pt_com.RoomItem{
			Roomid: *proto.Uint64(v.RoomID),
			Roomname: *proto.String(v.RoomName),
			Roomkind: *proto.Uint32(v.RoomKind),
		}
	}

	data := pt_logical.L2SLoginUserResult{
		Loginsocket: *proto.Uint32(loginsocket),
		Address: *proto.String(address),
		Code: *proto.Int32(retcode),
		Userid: *proto.Uint64(userid),
		Account: *proto.String(account),
		Password: *proto.String(password),
		Gold: *proto.Uint64(gold),
		Rooms: rm,
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LogicalServer SendUserLoginResult Error:", err.Error())
		return
	}

	o.SendPacket(logicalsocket, uint32(pt_login.LoginPid_LOGIN_LOGIN), bytes)
}

func (o *LogicalServer) SendUserRegisterResult(logicalsocket uint32, loginsocket uint32, retcode int32, account string, password string) {
	data := pt_logical.L2SRegisterUserResult{
		Loginsocket: *proto.Uint32(loginsocket),
		Code: *proto.Int32(retcode),
		Account: *proto.String(account),
		Password: *proto.String(password),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("LogicalServer SendUserRegisterResult Error:", err.Error())
		return
	}

	o.SendPacket(logicalsocket, uint32(pt_login.LoginPid_LOGIN_REGISTER), bytes)
}

func (o *LogicalServer) onUserLogin(socket uint32, bytes []byte) {
	var data pt_logical.S2LLoginUser
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LogicalServer onUserLogin Error: ", err.Error())
		return
	}

	fmt.Println("LogicalServer onUserLogin :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_LOGIN, socket, data.Loginsocket, data.Address, data.Account, data.Password)
}

func (o *LogicalServer) onUserRegister(socket uint32, bytes []byte) {
	var data pt_logical.S2LRegisterUser
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("LogicalServer onUserRegister Error: ", err.Error())
		return
	}

	fmt.Println("LogicalServer onUserRegister :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_REGISTER, socket, data.Loginsocket, data.Account, data.Password)
}