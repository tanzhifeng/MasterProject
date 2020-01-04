package include

import (
	"../../common"
	"../../common/prot/pt_com"
	"github.com/golang/protobuf/ptypes/any"
)

type ILogicalServer interface {
	SendRegisterNodeResult(logicalsocket uint32, code int32, nodetype uint32, extra *any.Any)
	SendUserLoginResult(logicalsocket uint32, loginsocket uint32, address string, retcode int32, userid uint64, account string, password string, gold uint64, rooms *map[uint64]*common.DataRoom)
	SendUserRegisterResult(logicalsocket uint32, loginsocket uint32, retcode int32, account string, password string)

	NotifyRoomAppend(logicalsocket uint32, roomid uint64, roomname string, roomkind uint32)
	NotifyRoomRemoved(logicalsocket uint32, roomid uint64)

	SendUserEnterRoomResult(logicalsocket uint32, loginsocket uint32, code int32, roomid uint64, users map[uint64]*pt_com.RoomUserItem, tables map[uint32]*pt_com.RoomTableItem)
	SendUserLeaveRoomResult(logicalsocket uint32, loginsocket uint32, code int32, roomid uint64)

	NotifyRoomUserAppend(logicalsocket uint32, loginsocket uint32, userid uint64, username string, gold uint64)
	NotifyRoomUserRemove(logicalsocket uint32, loginsocket uint32, userid uint64)
}