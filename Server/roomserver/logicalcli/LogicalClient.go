package logicalcli

import (
	"../../common"
	"../../common/network"
	"../../common/prot/pt_com"
	"../logic"
	"fmt"
	"sync"
	"time"
)

type LogicalClient struct {
	ip string
	port uint
	conn *network.ClientTCP
	vaild bool
	persist bool
	channelStart chan bool
	channelStop chan bool
	map4Handler map[uint32]network.HandlerS2C
}

var logicalC *LogicalClient
var logicalCOnce sync.Once

func GetLogicalClient() *LogicalClient {
	logicalCOnce.Do(func() {
		logicalC = &LogicalClient{}
		logicalC.init()
	})
	return logicalC
}

func (o *LogicalClient) init() {
	o.map4Handler = make(map[uint32]network.HandlerS2C)

	o.map4Handler[uint32(pt_com.ComPid_COM_REGISTER_NODE)] = o.onRegisterNodeResult
}

func (o *LogicalClient) Start(ip string, port uint, persist bool) bool {
	o.ip = ip
	o.port = port
	o.conn = new(network.ClientTCP)
	ok := o.conn.Start(o.ip, o.port)
	if !ok {
		fmt.Println("LogicalClient Start Failed !")
		return false
	}

	o.vaild = true
	o.persist = persist
	o.channelStart = make(chan bool)
	o.channelStop = make(chan bool)

	go o.HandleUpdate()

	defer close(o.channelStart)

	v, ok := <-o.channelStart

	if ok && v {
		fmt.Println("LogicalClient Start Successed !")
	} else {
		fmt.Println("LogicalClient Start Failed !")
	}

	return v
}

func (o *LogicalClient) Stop() {
	o.persist = false
	o.vaild = false

	o.conn.Stop()

	<-o.channelStop
}

func (o *LogicalClient) SendPacket(ptid uint32, content []byte) {
	if o.vaild {
		o.conn.SendClientPacket(ptid, content)
	}
}

func (o *LogicalClient) HandlerReconnect() {
	connected := !o.persist

	for !connected {
		time.Sleep(time.Second * 5)
		connected = o.Start(o.ip, o.port, o.persist)
	}
}

func (o *LogicalClient) HandleUpdate() {
	defer o.HandlerReconnect()
	defer close(o.channelStop)

	for {
		v, ok := o.conn.GetClientPacket()
		if !ok {
			return
		} else {
			switch v.Status {
			case network.Connected:
				fmt.Println("LogicalClient Connected !")
				o.SendRegisterNode()
				logic.GetLogical().AppendEvent(common.EVENT_LOGICAL_CLIENT_CONNECTED)
			case network.Losted:
				fmt.Println("LogicalClient Losted !")
				o.vaild = false
				logic.GetLogical().AppendEvent(common.EVENT_LOGICAL_CLIENT_LOSTED)
			case network.Dataed:
				fmt.Printf("LogicalClient Connection Dataed Ptid: %d !\n", v.Ptid)
				handler, ok := o.map4Handler[v.Ptid]
				if !ok {
					fmt.Println("LogicalClient Unmatch Ptid: ", v.Ptid)
				} else {
					handler(v.Content)
				}
			default:
				fmt.Printf("LogicalClient Unknow Status %d !\n", v.Status)
			}
		}
	}
}
