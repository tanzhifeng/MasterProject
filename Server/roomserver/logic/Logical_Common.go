package logic

import (
	"../global"
	"fmt"
)

func (o *Logical) LogicalClientConnected(params  ...interface{}) {
	fmt.Println(params...)
}

func (o *Logical) LogicalClientLosted(params  ...interface{}) {
	fmt.Println(params...)
}

func (o *Logical) RoomServerClientConnected(params ...interface{}) {
	fmt.Println(params...)

	var socket = params[0].(uint32)
	var address = params[1].(string)
	global.GetGlobalData().AddOnlineSocket(socket, address)
}

func (o *Logical) RoomServerClientLost(params ...interface{}) {
	//fmt.Println(params...)
	//
	//var socket = params[0].(uint32)
	//v := global.GetGlobalData().GetOnlineSocket(socket)
	//if v != nil && v.SocketType == common.SOCKET_USER {
	//	//玩家掉线
	//	component, ok := tools.GetComponentCollecter().GetComponent(common.LogicalClient)
	//	if !ok {
	//		return
	//	}
	//	logicalcli := component.(include.ILogicalClient)
	//	logicalcli.SendUserLosted(socket)
	//}
	//global.GetGlobalData().RemoveOnlineSocket(socket)
}