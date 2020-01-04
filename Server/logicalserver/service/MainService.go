package service

import (
	"fmt"

	"../../common"
	"../../common/tools"
	"../logic"
	"../logicalsvr"
	"../databasecli"
)

func Start(config map[string]interface{}) bool {
	tools.GetComponentCollecter().AddComponent(common.Scheduler, tools.GetScheduler())

	ok := logic.GetLogical().Start()
	if !ok {
		fmt.Println("Logical Start Failed !")
		defer Stop()
		return false
	}
	tools.GetComponentCollecter().AddComponent(common.Logical, logic.GetLogical())

	databaseip := config["databaseip"].(string)
	databaseport := uint(config["databaseport"].(float64))

	ok = databasecli.GetDatabaseClient().Start(databaseip, databaseport, true)
	if !ok {
		fmt.Println("DatabaseClient Start Failed !")
		defer Stop()
		return false
	}
	tools.GetComponentCollecter().AddComponent(common.DatabaseClient, databasecli.GetDatabaseClient())

	logicalip := config["logicalip"].(string)
	logicalport := uint(config["logicalport"].(float64))

	ok = logicalsvr.GetLogicalServer().Start(logicalip, logicalport)
	if !ok {
		fmt.Println("LogicalServer Start Failed !")
		defer Stop()
		return false
	}
	tools.GetComponentCollecter().AddComponent(common.LogicalServer, logicalsvr.GetLogicalServer())

	fmt.Println("LogicalService Start Successed !")

	return true
}

func Stop() {
	tools.GetComponentCollecter().StopComponent(common.Scheduler)
	tools.GetComponentCollecter().StopComponent(common.LogicalServer)
	tools.GetComponentCollecter().StopComponent(common.DatabaseClient)
	tools.GetComponentCollecter().StopComponent(common.Logical)

	fmt.Println("LogicalService Stop Successed !")
}
