package include

type IDatabaseServer interface {
	SendUserLoginResult(socket uint32, code int32, logicalsocket uint32, loginsocket uint32, address string, userid uint64, account string, password string, gold uint64)
	SendUserRegisterResult(socket uint32, code int32, logicalsocket uint32, loginsocket uint32, account string, password string)
}