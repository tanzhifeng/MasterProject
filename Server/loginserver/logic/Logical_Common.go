package logic

import (
	"../../common"
	"../../common/tools"
	"../global"
	"../include"
	"fmt"
	"time"
)

func (o *Logical) LogicalClientConnected(params  ...interface{}) {
	fmt.Println(params...)
}

func (o *Logical) LogicalClientLosted(params  ...interface{}) {
	fmt.Println(params...)
}

func (o *Logical) LoginServerClientConnected(params ...interface{}) {
	fmt.Println(params...)

	var socket = params[0].(uint32)
	var address = params[1].(string)
	global.GetGlobalData().AddOnlineSocket(socket, address)
}

func (o *Logical) LoginServerClientLost(params ...interface{}) {
	fmt.Println(params...)

	var socket = params[0].(uint32)
	v := global.GetGlobalData().GetOnlineSocket(socket)
	if v != nil && v.SocketType == common.SOCKET_USER {
		//玩家掉线
		component, ok := tools.GetComponentCollecter().GetComponent(common.LogicalClient)
		if !ok {
			return
		}
		logicalcli := component.(include.ILogicalClient)
		logicalcli.SendUserLosted(socket)
	}
	global.GetGlobalData().RemoveOnlineSocket(socket)
}

func (o *Logical) LoginServerClientCheckAlive(params ...interface{}) {
	component, ok := tools.GetComponentCollecter().GetComponent(common.LoginServer)
	if !ok {
		return
	}
	loginserver := component.(include.ILoginServer)

	nowTime := time.Now().Unix()
	sockets := global.GetGlobalData().GetSockets()

	for _, v := range *sockets {
		delta := nowTime - v.BeatsLast
		if delta > common.BeatsIntervalServer {
			loginserver.CloseClient(v.Socket)
		}
	}
}

func (o *Logical) LoginServerClientBeats(params ...interface{}) {
	fmt.Println(params...)

	var socket = params[0].(uint32)
	global.GetGlobalData().UpdateSocketBeats(socket)
}