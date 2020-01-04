package service

import (
	"fmt"

	"../../common"
	"../../common/tools"
	"../logic"
	"../logicalcli"
	"../loginsvr"
)

func Start(config map[string]interface{}) bool {
	tools.GetComponentCollecter().AddComponent(common.Scheduler, tools.GetScheduler())

	ok := logic.GetLogical().Start()
	if !ok {
		defer Stop()
		return false
	}
	tools.GetComponentCollecter().AddComponent(common.Logical, logic.GetLogical())

	logicalip := config["logicalip"].(string)
	logicalport := uint(config["logicalport"].(float64))

	ok = logicalcli.GetLogicalClient().Start(logicalip, logicalport, true)
	if !ok {
		defer Stop()
		return false
	}
	tools.GetComponentCollecter().AddComponent(common.LogicalClient, logicalcli.GetLogicalClient())

	loginip := config["loginip"].(string)
	loginport := uint(config["loginport"].(float64))

	ok = loginsvr.GetLoginServer().Start(loginip, loginport)
	if !ok {
		defer Stop()
		return false
	}
	tools.GetComponentCollecter().AddComponent(common.LoginServer, loginsvr.GetLoginServer())

	fmt.Println("LoginService Start Successed !")

	return true
}

func Stop() {
	tools.GetComponentCollecter().StopComponent(common.Scheduler)
	tools.GetComponentCollecter().StopComponent(common.LoginServer)
	tools.GetComponentCollecter().StopComponent(common.LogicalClient)
	tools.GetComponentCollecter().StopComponent(common.Logical)

	fmt.Println("LoginService Stop Successed !")
}
