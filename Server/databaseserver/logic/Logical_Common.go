package logic

import (
	"../../common/mysql"
	"../mysqlcli"
	"fmt"
)

func (o *Logical) LogicalServerConnected(params ...interface{}) {
	fmt.Println(params...)
}

func (o *Logical) LogicalUserLosted(params ...interface{}) {
	fmt.Println(params...)

	userid := params[1].(uint64)
	gold := params[2].(uint64)

	operate := &mysql.Operate{
		OperateType: mysql.OperateTransaction,
		OperateList: []*mysql.OperateSingle{
			{
				OperateID: mysqlcli.DB_USER_SAVE,
				OperateArgs: []interface{}{gold, userid},
			},
		},
		OperateOptions: params,
		Handler: o.DBUserSave,
	}
	mysqlcli.GetMySQLClient().Execute(userid, operate)
}

func (o *Logical) DBUserSave(result *mysql.OperateResult) {
	userid := result.OperateOptions[1].(uint64)

	if result.Success {
		fmt.Printf("[%d] User Data Save Successed !\n", userid)
	} else {
		fmt.Printf("[%d] User Data Save Failed !\n", userid)
	}

	mysqlcli.GetMySQLClient().DeleteClientChannel(userid)
}