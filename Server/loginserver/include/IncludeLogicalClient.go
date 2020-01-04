package include

type ILogicalClient interface {
	SendRegisterNode()
	SendUserLosted(loginsocket uint32)
	SendUserLogin(loginsocket uint32, address string, account string, password string)
	SendUserRegister(loginsocket uint32, account string, password string)

	SendUserEnterRoom(loginsocket uint32, roomid uint64)
	SendUserLeaveRoom(loginsocket uint32)
}