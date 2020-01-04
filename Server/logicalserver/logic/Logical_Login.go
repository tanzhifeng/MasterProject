package logic

import (
	"../../common"
	"../../common/prot/pt_login"
	"../../common/tools"
	"../global"
	"../include"
	"fmt"
)

func (o *Logical) LogicalUserLogin(params ...interface{}) {
	fmt.Println(params...)

	component, ok := tools.GetComponentCollecter().GetComponent(common.DatabaseClient)
	if !ok {
		return
	}

	logicalsocket := params[0].(uint32)
	loginsocket := params[1].(uint32)
	address := params[2].(string)
	account := params[3].(string)
	password := params[4].(string)

	databasecli := component.(include.IDatabaseClient)
	databasecli.SendUserLogin(logicalsocket, loginsocket, address, account, password)
}

func (o *Logical) LogicalUserLoginResult(params ...interface{}) {
	fmt.Println(params...)

	component, ok := tools.GetComponentCollecter().GetComponent(common.LogicalServer)
	if !ok {
		return
	}

	logicalsocket := params[0].(uint32)
	loginsocket := params[1].(uint32)
	address := params[2].(string)
	code := params[3].(int32)
	userid := params[4].(uint64)
	account := params[5].(string)
	password := params[6].(string)
	gold := params[7].(uint64)
	rooms := global.GetGlobalData().GetRooms()

	if code == int32(pt_login.LoginCode_C_SUCCESS) {
		if global.GetGlobalData().IsUserOnlineByUserId(userid) {
			//重复登录
			code = int32(pt_login.LoginCode_C_ALREADY_ONLINE)
		} else {
			//登录成功
			global.GetGlobalData().AddUser(logicalsocket, loginsocket, address, userid, account, password, gold)
		}
	}
	logicalserver := component.(include.ILogicalServer)
	logicalserver.SendUserLoginResult(logicalsocket, loginsocket, address, code, userid, account,password, gold, rooms)
}

func (o *Logical) LogicalUserRegister(params ...interface{}) {
	fmt.Println(params...)

	component, ok := tools.GetComponentCollecter().GetComponent(common.DatabaseClient)
	if !ok {
		return
	}

	logicalsocket := params[0].(uint32)
	loginsocket := params[1].(uint32)
	account := params[2].(string)
	password := params[3].(string)

	databasecli := component.(include.IDatabaseClient)
	databasecli.SendUserRegister(logicalsocket, loginsocket, account, password)
}

func (o *Logical) LogicalUserRegisterResult(params ...interface{}) {
	fmt.Println(params...)

	component, ok := tools.GetComponentCollecter().GetComponent(common.LogicalServer)
	if !ok {
		return
	}

	logicalsocket := params[0].(uint32)
	loginsocket := params[1].(uint32)
	code := params[2].(int32)
	account := params[3].(string)
	password := params[4].(string)

	logicalserver := component.(include.ILogicalServer)
	logicalserver.SendUserRegisterResult(logicalsocket, loginsocket, code, account,password)
}