package service

import (
	"fmt"

	"../../common"
	"../../common/tools"
	"../logic"
	"../logicalcli"
	"../roomsvr"
)

func Start(config map[string]interface{}) bool {
	tools.GetComponentCollecter().AddComponent(common.Scheduler, tools.GetScheduler())

	ok := logic.GetLogical().Start(&config)
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

	roomip := config["roomip"].(string)
	roomport := uint(config["roomport"].(float64))

	ok = roomsvr.GetRoomServer().Start(roomip, roomport)
	if !ok {
		defer Stop()
		return false
	}
	tools.GetComponentCollecter().AddComponent(common.RoomServer, roomsvr.GetRoomServer())

	fmt.Println("RoomService Start Successed !")

	return true
}

func Stop() {
	tools.GetComponentCollecter().StopComponent(common.Scheduler)
	tools.GetComponentCollecter().StopComponent(common.RoomServer)
	tools.GetComponentCollecter().StopComponent(common.LogicalClient)
	tools.GetComponentCollecter().StopComponent(common.Logical)

	fmt.Println("RoomService Stop Successed !")
}
