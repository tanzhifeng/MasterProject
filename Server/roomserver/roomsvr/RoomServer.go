package roomsvr

import (
	"../../common"
	"../../common/network"
	"../../common/prot/pt_login"
	"../logic"
	"fmt"
	"sync"
)

type RoomServer struct {
	ip string
	port uint
	conn *network.ServerTCP
	map4Handler map[uint32]network.HandlerC2S
	channelStop chan bool
	scheduleBeats uint
}

var roomS *RoomServer
var roomSOnce sync.Once

func GetRoomServer() *RoomServer {
	roomSOnce.Do(func() {
		roomS = &RoomServer{}
		roomS.init()
	})
	return roomS
}

func (o *RoomServer) init() {
	o.map4Handler = make(map[uint32]network.HandlerC2S)

	o.map4Handler[uint32(pt_login.LoginPid_LOGIN_LOST)] = o.onLoginClientLost
	o.map4Handler[uint32(pt_login.LoginPid_LOGIN_BEATS)] = o.onLoginClientBeats
}

func (o *RoomServer) Start(ip string, port uint) bool {
	o.ip = ip
	o.port = port
	o.conn = new(network.ServerTCP)
	ok := o.conn.Start(o.ip, o.port)
	if !ok {
		fmt.Println("RoomServer Start Failed !")
		return false
	}

	o.channelStop = make(chan bool)

	go o.HandleUpdate()

	fmt.Println("RoomServer Start Successed !")

	return true
}

//Stop ...
func (o *RoomServer) Stop() {
	if o.conn != nil {
		o.conn.Stop()

		<-o.channelStop
	}
}

//SendPacket ...
func (o *RoomServer) SendPacket(socket uint32, ptid uint32, content []byte) {
	o.conn.SendServerPacket(socket, ptid, content)
}

//CloseClient ...
func (o *RoomServer) CloseClient(socket uint32) {
	o.conn.CloseServerClient(socket)
}

func (o *RoomServer) HandleUpdate() {
	defer close(o.channelStop)

	for {
		v, ok := o.conn.GetServerPacket()
		if !ok {
			return
		} else {
			switch v.Status {
			case network.Connected:
				fmt.Printf("RoomClient Connected SocketId:%d !\n", v.Socket)
				logic.GetLogical().AppendEvent(common.EVENT_ROOM_SERVER_CLIENT_CONNECTED, v.Socket, o.conn.GetClientAddr(v.Socket))
			case network.Losted:
				fmt.Printf("RoomClient Losted SocketId:%d !\n", v.Socket)
				logic.GetLogical().AppendEvent(common.EVENT_ROOM_SERVER_CLIENT_LOSTED, v.Socket)
			case network.Dataed:
				fmt.Printf("RoomClient Connection Dataed SocketId:%d Ptid: %d !\n", v.Socket, v.Ptid)
				handler, ok := o.map4Handler[v.Ptid]
				if !ok {
					fmt.Println("RoomClient Unmatch Ptid: ", v.Ptid)
				} else {
					handler(v.Socket, v.Content)
				}
			default:
				fmt.Printf("RoomClient Unknow Status SocketId:%d Status:%d !\n", v.Socket, v.Status)
			}
		}
	}
}