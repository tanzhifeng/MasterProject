package databasesvr

import (
	"../../common"
	"../../common/prot/pt_database"
	"../logic"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func (o *DatabaseServer) onUserDataSave(socket uint32, bytes []byte) {
	var data pt_database.L2DUserDataSave
	err := proto.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("DatabaseServer onUserDataSave Error: ", err.Error())
		return
	}

	fmt.Println("DatabaseServer onUserDataSave :", data)

	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_LOSTED, socket, data.Userid, data.Gold)
}