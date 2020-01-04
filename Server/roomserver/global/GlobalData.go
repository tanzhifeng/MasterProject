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
}

//获取所有在线连接
func (o *GlobalData) GetSockets() *map[uint32]*common.DataSocket {
	return &o.MapSocket
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