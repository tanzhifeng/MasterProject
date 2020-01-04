package logicalsvr

import (
	"fmt"
	"sync"

	"../../common"
	"../../common/network"
	"../../common/prot/pt_com"
	"../../common/prot/pt_login"
	"../logic"
)

type LogicalServer struct {
	ip string
	port uint
	conn *network.ServerTCP
	map4Handler map[uint32]network.HandlerC2S
	channelStop chan bool
}

var logicalS *LogicalServer
var logicalSOnce sync.Once

func GetLogicalServer() *LogicalServer {
	logicalSOnce.Do(func() {
		logicalS = &LogicalServer{}
		logicalS.init()
	})
	return logicalS
}

func (o *LogicalServer) init() {
	o.map4Handler = make(map[uint32]network.HandlerC2S)

	o.map4Handler[uint32(pt_com.ComPid_COM_REGISTER_NODE)] = o.onRegisterNode
	o.map4Handler[uint32(pt_login.LoginPid_LOGIN_LOST)] = o.onUserLosted
	o.map4Handler[uint32(pt_login.LoginPid_LOGIN_LOGIN)] = o.onUserLogin
	o.map4Handler[uint32(pt_login.LoginPid_LOGIN_REGISTER)] = o.onUserRegister
	o.map4Handler[uint32(pt_login.LoginPid_LOGIN_ROOM_USER_ENTER)] = o.onUserEnterRoom
	o.map4Handler[uint32(pt_login.LoginPid_LOGIN_ROOM_USER_LEAVE)] = o.onUserLeaveRoom
}

func (o *LogicalServer) Start(ip string, port uint) bool {
	o.ip = ip
	o.port = port
	o.conn = new(network.ServerTCP)
	ok := o.conn.Start(o.ip, o.port)
	if !ok {
		fmt.Println("LogicalServer Start Failed !")
		return false
	}

	o.channelStop = make(chan bool)

	go o.HandleUpdate()

	fmt.Println("LogicalServer Start Successed !")

	return true
}

//Stop ...
func (o *LogicalServer) Stop() {
	if o.conn != nil {
		o.conn.Stop()

		<-o.channelStop
	}
}

//SendPacket ...
func (o *LogicalServer) SendPacket(socket uint32, ptid uint32, content []byte) {
	o.conn.SendServerPacket(socket, ptid, content)
}

//CloseClient ...
func (o *LogicalServer) CloseClient(socket uint32) {
	o.conn.CloseServerClient(socket)
}

func (o *LogicalServer) HandleUpdate() {
	defer close(o.channelStop)

	for {
		v, ok := o.conn.GetServerPacket()
		if !ok {
			return
		} else {
			switch v.Status {
			case network.Connected:
				fmt.Printf("LogicalClient Connected SocketId:%d !\n", v.Socket)
				logic.GetLogical().AppendEvent(common.EVENT_LOGICAL_SERVER_CLIENT_CONNECTED, v.Socket, o.conn.GetClientAddr(v.Socket))
			case network.Losted:
				fmt.Printf("LogicalClient Losted SocketId:%d !\n", v.Socket)
				logic.GetLogical().AppendEvent(common.EVENT_LOGICAL_SERVER_CLIENT_LOSTED, v.Socket)
			case network.Dataed:
				fmt.Printf("LogicalClient Connection Dataed SocketId:%d Ptid: %d !\n", v.Socket, v.Ptid)
				handler, ok := o.map4Handler[v.Ptid]
				if !ok {
					fmt.Println("LogicalClient Unmatch Ptid: ", v.Ptid)
				} else {
					handler(v.Socket, v.Content)
				}
			default:
				fmt.Printf("LogicalClient Unknow Status SocketId:%d Status:%d !\n", v.Socket, v.Status)
			}
		}
	}
}