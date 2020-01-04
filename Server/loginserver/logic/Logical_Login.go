package logic

import (
	"../../common"
	"../../common/prot/pt_com"
	"../../common/prot/pt_login"
	"../../common/tools"
	"../global"
	"../include"
	"fmt"
)

func (o *Logical) LoginServerUserLogin(params ...interface{}) {
	fmt.Println(params...)

	loginsocket := params[0].(uint32)
	account := params[1].(string)
	password := params[2].(string)

	if global.GetGlobalData().IsOnlineUser(loginsocket) || global.GetGlobalData().IsOnlineAccount(account) {
		component, ok := tools.GetComponentCollecter().GetComponent(common.LoginServer)
		if !ok {
			return
		}
		loginserver := component.(include.ILoginServer)
		loginserver.SendUserLoginResult(loginsocket, int32(pt_login.LoginCode_C_ALREADY_ONLINE), 0, account, password, 0, nil)
	} else {
		component, ok := tools.GetComponentCollecter().GetComponent(common.LogicalClient)
		if !ok {
			return
		}

		v := global.GetGlobalData().GetOnlineSocket(loginsocket)
		if v == nil {
			return
		}

		logicalcli := component.(include.ILogicalClient)
		logicalcli.SendUserLogin(loginsocket, v.Address, account, password)
	}
}

func (o *Logical) LoginServerUserLoginResult(params ...interface{}) {
	fmt.Println(params...)

	component, ok := tools.GetComponentCollecter().GetComponent(common.LoginServer)
	if !ok {
		return
	}

	loginsocket := params[0].(uint32)
	code := params[1].(int32)
	userid := params[2].(uint64)
	account := params[3].(string)
	password := params[4].(string)
	gold := params[5].(uint64)
	rooms := params[6].(map[uint64]*pt_com.RoomItem)

	if code == int32(pt_login.LoginCode_C_SUCCESS) {
		global.GetGlobalData().SetUserSocket(loginsocket, account)
	}

	loginserver := component.(include.ILoginServer)
	loginserver.SendUserLoginResult(loginsocket, code, userid, account, password, gold, rooms)
}

func (o *Logical) LoginServerUserRegister(params ...interface{}) {
	fmt.Println(params...)

	component, ok := tools.GetComponentCollecter().GetComponent(common.LogicalClient)
	if !ok {
		return
	}

	loginsocket := params[0].(uint32)
	account := params[1].(string)
	password := params[2].(string)

	logicalcli := component.(include.ILogicalClient)
	logicalcli.SendUserRegister(loginsocket, account, password)
}

func (o *Logical) LoginServerUserRegisterResult(params ...interface{}) {
	fmt.Println(params...)

	component, ok := tools.GetComponentCollecter().GetComponent(common.LoginServer)
	if !ok {
		return
	}

	loginsocket := params[0].(uint32)
	code := params[1].(int32)
	account := params[2].(string)
	password := params[3].(string)

	loginserver := component.(include.ILoginServer)
	loginserver.SendUserRegisterResult(loginsocket, code, account, password)
}
