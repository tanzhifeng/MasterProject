package global

import (
	"sync"
	"time"

	"../../common"
)

var globaldata *GlobalData
var globaldataOnce sync.Once

type GlobalData struct {
	MapSocket map[uint32]*common.DataSocket //在线连接
	MapAccount map[string]uint32 //在线帐号记录
}

func GetGlobalData() *GlobalData {
	globaldataOnce.Do(func() {
		globaldata = &GlobalData{}
		globaldata.init()
	})
	return globaldata
}

func (o *GlobalData) init() {
	o.MapSocket = make(map[uint32]*common.DataSocket)
	o.MapAccount = make(map[string]uint32)
}

//获取所有在线连接
func (o *GlobalData) GetSockets() *map[uint32]*common.DataSocket {
	return &o.MapSocket
}

func (o *GlobalData) GetAccounts() *map[string]uint32 {
	return &o.MapAccount
}

//新增在线连接
func (o *GlobalData) AddOnlineSocket(socket uint32, address string) {
	o.MapSocket[socket] = &common.DataSocket{
		Socket: socket,
		SocketType: common.SOCKET_UNKNOW,
		Address: address,
		BeatsLast: time.Now().Unix(),
	}
}

//获取在线连接数据
func (o *GlobalData) GetOnlineSocket(socket uint32) *common.DataSocket {
	return o.MapSocket[socket]
}

//删除在线连接
func (o *GlobalData) RemoveOnlineSocket(socket uint32) {
	v, ok := o.MapSocket[socket]
	if ok {
		switch v.SocketType {
		case common.SOCKET_USER:
			delete(o.MapAccount, v.Extra)
		}
		delete(o.MapSocket, socket)
	}
}

//刷新连接心跳记录
func (o *GlobalData) UpdateSocketBeats(socket uint32) {
	dataOnline, ok := o.MapSocket[socket]
	if ok {
		dataOnline.BeatsLast = time.Now().Unix()
	}
}

//设置为用户连接
func (o *GlobalData) SetUserSocket(socket uint32, account string) {
	dataOnline, ok := o.MapSocket[socket]
	if ok {
		dataOnline.SocketType = common.SOCKET_USER
		dataOnline.Extra = account
		dataOnline.BeatsLast = time.Now().Unix()

		o.MapAccount[account] = socket
	}
}

//判断帐号是否已登录
func (o *GlobalData) IsOnlineAccount(account string) bool {
	_, ok := o.MapAccount[account]

	return ok
}

//判断Socket是否已登录
func (o *GlobalData) IsOnlineUser(socket uint32) bool {
	online, ok := o.MapSocket[socket]
	if ok {
		return online.SocketType == common.SOCKET_USER
	} else {
		return false
	}
}