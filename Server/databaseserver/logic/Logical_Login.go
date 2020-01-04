package logic

import (
	"../../common"
	"../../common/mysql"
	"../../common/prot/pt_login"
	"../../common/tools"
	"../include"
	"../mysqlcli"
	"fmt"
)

func (o *Logical) LogicalUserLogin(params ...interface{}) {
	fmt.Println(params...)

	account := params[4].(string)
	password := params[5].(string)

	operate := &mysql.Operate{
		OperateType: mysql.OperateCommand,
		OperateList: []*mysql.OperateSingle{
			{
				OperateID: mysqlcli.DB_USER_LOGIN,
				OperateArgs: []interface{}{account, password},
			},
		},
		OperateOptions: params,
		Handler: o.DBUserLogin,
	}
	mysqlcli.GetMySQLClient().Execute(0, operate)
}

func (o *Logical) DBUserLogin(result *mysql.OperateResult) {
	component, ok := tools.GetComponentCollecter().GetComponent(common.DatabaseServer)
	if !ok {
		return
	}

	databasesvr := component.(include.IDatabaseServer)

	retSingle := result.Results[0]
	success := retSingle.Success && len(retSingle.QueryRows) > 0
	socket := result.OperateOptions[0].(uint32)
	logicalsocket := result.OperateOptions[1].(uint32)
	loginsocket := result.OperateOptions[2].(uint32)
	address := result.OperateOptions[3].(string)
	account := result.OperateOptions[4].(string)
	password := result.OperateOptions[5].(string)

	if success {
		fields := *(retSingle.QueryRows[0].(*[]interface{}))
		userid := *(fields[0].(*uint64))
		gold := *(fields[3].(*uint64))

		mysqlcli.GetMySQLClient().CreateClientChannel(userid)
		databasesvr.SendUserLoginResult(socket, int32(pt_login.LoginCode_C_SUCCESS), logicalsocket, loginsocket, address, userid, account, password, gold)
	} else {
		databasesvr.SendUserLoginResult(socket, int32(pt_login.LoginCode_C_FAILED), logicalsocket, loginsocket, address, 0, "", "", 0)
	}
}

func (o *Logical) LogicalUserRegister(params ...interface{}) {
	fmt.Println(params...)

	account := params[3].(string)
	password := params[4].(string)

	operate := &mysql.Operate{
		OperateType: mysql.OperateTransaction,
		OperateList: []*mysql.OperateSingle{
			{
				OperateID: mysqlcli.DB_USER_REGISTER,
				OperateArgs: []interface{}{account, password},
			},
			{
				OperateID: mysqlcli.DB_USER_SELECT_LAST,
			},
		},
		OperateOptions: params,
		Handler: o.DBUserRegister,
	}
	mysqlcli.GetMySQLClient().Execute(0, operate)
}

func (o *Logical) DBUserRegister(result *mysql.OperateResult) {
	component, ok := tools.GetComponentCollecter().GetComponent(common.DatabaseServer)
	if !ok {
		return
	}

	databasesvr := component.(include.IDatabaseServer)

	success := result.Success
	socket := result.OperateOptions[0].(uint32)
	logicalsocket := result.OperateOptions[1].(uint32)
	loginsocket := result.OperateOptions[2].(uint32)
	account := result.OperateOptions[3].(string)
	password := result.OperateOptions[4].(string)

	if success {
		databasesvr.SendUserRegisterResult(socket, int32(pt_login.LoginCode_C_SUCCESS), logicalsocket, loginsocket, account, password)
	} else {
		databasesvr.SendUserRegisterResult(socket, int32(pt_login.LoginCode_C_FAILED), logicalsocket, loginsocket, "", "")
	}
}