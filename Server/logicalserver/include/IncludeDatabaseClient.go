package include

type IDatabaseClient interface {
	SendUserDataSave(userid uint64, gold uint64)
	SendUserLogin(logicalsocket uint32, loginsocket uint32, address string, account string, password string)
	SendUserRegister(logicalsocket uint32, loginsocket uint32, account string, password string)
}