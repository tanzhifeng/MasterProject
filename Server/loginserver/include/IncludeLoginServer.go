package include

import (
	"../../common/prot/pt_com"
)

type ILoginServer interface {
	CloseClient(socket uint32)

	SendUserLoginResult(loginsocket uint32, code int32, userid uint64, account string, password string, gold uint64, rooms map[uint64]*pt_com.RoomItem)
	SendUserRegisterResult(loginsocket uint32, code int32, account string, password string)

	NotifyRoomAppend(loginsocket uint32, roomid uint64, roomname string, roomkind uint32)
	NotifyRoomRemove(loginsocket uint32, roomid uint64)

	SendUserEnterRoomResult(loginsocket uint32, code int32, roomid uint64, users map[uint64]*pt_com.RoomUserItem, tables map[uint32]*pt_com.RoomTableItem)
	SendUserLeaveRoomResult(loginsocket uint32, code int32, roomid uint64)

	NotifyRoomUserAppend(loginsocket uint32, userid uint64, username string, gold uint64, tableid uint32, chairid uint32)
	NotifyRoomUserRemove(loginsocket uint32, userid uint64)
}