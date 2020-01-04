package databasecli

import (
	"../../common/prot/pt_database"
	"../../common/prot/pt_login"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func (o *DatabaseClient) SendUserDataSave(userid uint64, gold uint64) {
	data := pt_database.L2DUserDataSave{
		Userid: *proto.Uint64(userid),
		Gold: *proto.Uint64(gold),
	}

	bytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("DatabaseClient SendUserDataSave Error:", err.Error())
		return
	}

	o.SendPacket(uint32(pt_login.LoginPid_LOGIN_LOST), bytes)
}