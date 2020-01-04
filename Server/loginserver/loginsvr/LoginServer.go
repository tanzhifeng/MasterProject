package loginsvr

import (
	"../../common"
	"../../common/network"
	"../../common/prot/pt_login"
	"../logic"
	"fmt"
	"sync"
)

type LoginServer struct {
	ip string
	port uint
	conn *network.ServerTCP
	map4Handler map[uint32]network.HandlerC2S
	channelStop chan bool
	scheduleBeats uint
}

var loginS *LoginServer
var loginSOnce sync.Once

func GetLoginServer() *LoginServer {
	loginSOnce.Do(func() {
		loginS = &LoginServer{}
		loginS.init()
	})
	return loginS
}

func (o *LoginServer) init() {
	o.map4Handler = make(map[uint32]network.HandlerC2S)

	o.map4Handler[uint32(pt_login.LoginPid_LOGIN_LOST)] = o.onLoginClientLost
	o.map4Handler[uint32(pt_login.LoginPid_LOGIN_BEATS)] = o.onLoginClientBeats

	o.map4Handler[uint32(pt_login.LoginPid_LOGIN_LOGIN)] = o.onUserLogin
	o.map4Handler[uint32(pt_login.LoginPid_LOGIN_REGISTER)] = o.onUserRegister
	o.map4Handler[uint32(pt_login.LoginPid_LOGIN_ROOM_USER_ENTER)] = o.onUserEnterRoom
	o.map4Handler[uint32(pt_login.LoginPid_LOGIN_ROOM_USER_LEAVE)] = o.onUserLeaveRoom

	//o.MapHandler[uint32(pt_gate.GatePt_GT_ROOM_USER_ENTER)] = o.On_GT_RoomUserEnter
	//o.MapHandler[uint32(pt_gate.GatePt_GT_ROOM_USER_LEAVE)] = o.On_GT_RoomUserLeave
	//o.MapHandler[uint32(pt_gate.GatePt_GT_ROOM_USER_SITDOWN)] = o.On_GT_RoomUserSitdown
	//o.MapHandler[uint32(pt_gate.GatePt_GT_ROOM_USER_STANDUP)] = o.On_GT_RoomUserStandup
	//
	//o.MapHandler[uint32(pt_room.RoomPt_RM_ROOM_LOGIN)] = o.On_RM_Room_Login
	//o.MapHandler[uint32(pt_room.RoomPt_RM_TABLE_LOGIN)] = o.On_RM_Table_Login
	//o.MapHandler[uint32(pt_room.RoomPt_RM_TABLE_LOST)] = o.On_RM_Table_Lost
}

func (o *LoginServer) Start(ip string, port uint) bool {
	o.ip = ip
	o.port = port
	o.conn = new(network.ServerTCP)
	ok := o.conn.Start(o.ip, o.port)
	if !ok {
		fmt.Println("LoginServer Start Failed !")
		return false
	}

	o.channelStop = make(chan bool)

	go o.HandleUpdate()

	o.checkBeatsStart()

	fmt.Println("LoginServer Start Successed !")

	return true
}

//Stop ...
func (o *LoginServer) Stop() {
	o.checkBeatsStop()

	if o.conn != nil {
		o.conn.Stop()

		<-o.channelStop
	}
}

//SendPacket ...
func (o *LoginServer) SendPacket(socket uint32, ptid uint32, content []byte) {
	o.conn.SendServerPacket(socket, ptid, content)
}

//CloseClient ...
func (o *LoginServer) CloseClient(socket uint32) {
	o.conn.CloseServerClient(socket)
}

func (o *LoginServer) HandleUpdate() {
	defer close(o.channelStop)

	for {
		v, ok := o.conn.GetServerPacket()
		if !ok {
			return
		} else {
			switch v.Status {
			case network.Connected:
				fmt.Printf("LoginClient Connected SocketId:%d !\n", v.Socket)
				logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_CONNECTED, v.Socket, o.conn.GetClientAddr(v.Socket))
			case network.Losted:
				fmt.Printf("LoginClient Losted SocketId:%d !\n", v.Socket)
				logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_LOSTED, v.Socket)
			case network.Dataed:
				fmt.Printf("LoginClient Connection Dataed SocketId:%d Ptid: %d !\n", v.Socket, v.Ptid)
				handler, ok := o.map4Handler[v.Ptid]
				if !ok {
					fmt.Println("LoginClient Unmatch Ptid: ", v.Ptid)
				} else {
					handler(v.Socket, v.Content)
				}
			default:
				fmt.Printf("LoginClient Unknow Status SocketId:%d Status:%d !\n", v.Socket, v.Status)
			}
		}
	}
}