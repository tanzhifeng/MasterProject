package databasecli

import (
	"../../common"
	"../../common/network"
	"../../common/prot/pt_login"
	"../logic"
	"fmt"
	"sync"
	"time"
)

type DatabaseClient struct {
	ip string
	port uint
	conn *network.ClientTCP
	vaild bool
	persist bool
	channelStop chan bool
	map4Handler map[uint32]network.HandlerS2C
}

var databaseC *DatabaseClient
var databaseCOnce sync.Once

func GetDatabaseClient() *DatabaseClient {
	databaseCOnce.Do(func() {
		databaseC = &DatabaseClient{}
		databaseC.init()
	})
	return databaseC
}

func (o *DatabaseClient) init() {
	o.map4Handler = make(map[uint32]network.HandlerS2C)

	o.map4Handler[uint32(pt_login.LoginPid_LOGIN_LOGIN)] = o.onUserLoginResult
	o.map4Handler[uint32(pt_login.LoginPid_LOGIN_REGISTER)] = o.onUserRegisterResult
	//o.MapHandler[uint32(pt_gate.GatePt_GT_REGISTER)] = o.On_GT_RegisterResult
	//
	//o.MapHandler[uint32(pt_room.RoomPt_RM_ROOM_LOGIN)] = o.On_RM_Room_LoginResult
	//o.MapHandler[uint32(pt_room.RoomPt_RM_TABLE_LOGIN)] = o.On_RM_Table_LoginResult
}

func (o *DatabaseClient) Start(ip string, port uint, persist bool) bool {
	o.ip = ip
	o.port = port
	o.conn = new(network.ClientTCP)
	ok := o.conn.Start(o.ip, o.port)
	if !ok {
		fmt.Println("DatabaseClient Start Failed !")
		return false
	}

	o.vaild = true
	o.persist = persist
	o.channelStop = make(chan bool)

	go o.HandleUpdate()

	fmt.Println("DatabaseClient Start Successed !")

	return true
}

func (o *DatabaseClient) Stop() {
	o.persist = false
	o.vaild = false

	o.conn.Stop()

	<-o.channelStop
}

func (o *DatabaseClient) SendPacket(ptid uint32, content []byte) {
	if o.vaild {
		o.conn.SendClientPacket(ptid, content)
	}
}

func (o *DatabaseClient) HandlerReconnect() {
	connected := !o.persist

	for !connected {
		time.Sleep(time.Second * 5)
		connected = o.Start(o.ip, o.port, o.persist)
	}
}

func (o *DatabaseClient) HandleUpdate() {
	defer o.HandlerReconnect()
	defer close(o.channelStop)

	for {
		v, ok := o.conn.GetClientPacket()
		if !ok {
			return
		} else {
			switch v.Status {
			case network.Connected:
				fmt.Println("DatabaseClient Connected !")
				logic.GetLogical().AppendEvent(common.EVENT_DATABASE_CLIENT_CONNECTED)
			case network.Losted:
				fmt.Println("DatabaseClient Losted !")
				o.vaild = false
				logic.GetLogical().AppendEvent(common.EVENT_DATABASE_CLIENT_LOSTED)
			case network.Dataed:
				fmt.Printf("DatabaseClient Connection Dataed Ptid: %d !\n", v.Ptid)
				handler, ok := o.map4Handler[v.Ptid]
				if !ok {
					fmt.Println("DatabaseClient Unmatch Ptid: ", v.Ptid)
				} else {
					handler(v.Content)
				}
			default:
				fmt.Printf("DatabaseClient Unknow Status %d !\n", v.Status)
			}
		}
	}
}