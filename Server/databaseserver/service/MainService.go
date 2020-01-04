package service

import (
	"fmt"

	"../../common"
	"../../common/tools"
	"../logic"
	"../databasesvr"
	"../mysqlcli"
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

	hostdb := config["hostdb"].(string)
	portdb := uint(config["portdb"].(float64))

	account := config["account"].(string)
	password := config["password"].(string)
	dbname := config["dbname"].(string)

	maxOpenConn := int(config["maxOpenConn"].(float64))
	maxIdleConn := int(config["maxIdleConn"].(float64))

	ok = mysqlcli.GetMySQLClient().Start(hostdb, portdb, account, password, dbname, maxOpenConn, maxIdleConn)
	if !ok {
		fmt.Println("MySQLClient Start Failed !")
		defer Stop()
		return false
	}
	mysqlcli.GetMySQLClient().CreateClientChannel(0)
	tools.GetComponentCollecter().AddComponent(common.MySQLClient, mysqlcli.GetMySQLClient())

	databaseip := config["databaseip"].(string)
	databaseport := uint(config["databaseport"].(float64))

	ok = databasesvr.GetDatabaseServer().Start(databaseip, databaseport)
	if !ok {
		fmt.Println("DatabaseServer Start Failed !")
		defer Stop()
		return false
	}
	tools.GetComponentCollecter().AddComponent(common.DatabaseServer, databasesvr.GetDatabaseServer())

	fmt.Println("DatabaseService Start Successed !")

	return true
}

func Stop() {
	tools.GetComponentCollecter().StopComponent(common.Scheduler)
	tools.GetComponentCollecter().StopComponent(common.DatabaseServer)
	tools.GetComponentCollecter().StopComponent(common.MySQLClient)
	tools.GetComponentCollecter().StopComponent(common.Logical)

	fmt.Println("DatabaseService Stop Successed !")
}
