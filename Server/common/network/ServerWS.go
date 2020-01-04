package network

import (
	"../../common"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"strings"
	"time"
)

type WSClient struct {
	socket       uint32
	conn         *websocket.Conn
	channelStop chan bool
	channelWrite chan *PacketWS
}

type ServerWS struct {
	ip          string
	port        uint
	srv			*http.Server
	channelStop chan bool
	channelRead chan *PacketWS
	connections map[uint32]*WSClient
}

var WSAmout uint32

func WSAutoAmount() uint32 {
	WSAmout++

	return WSAmout
}

func (o *ServerWS) handlerServerRead(socket uint32) bool {
	wsc, ok := o.connections[socket]
	if !ok {
		return false
	}

	wsc.conn.SetReadDeadline(time.Now().Add(time.Millisecond * 5))

	pkg := new(map[string]interface{})
	err := websocket.JSON.Receive(wsc.conn, pkg)
	if err != nil {
		errstr := err.Error()
		if strings.Contains(errstr, "i/o timeout") {
			return true
		} else {
			fmt.Println("Server Content Read Error ", err.Error(), wsc.socket)
			o.channelRead <- &PacketWS{Status: Losted, Socket: socket}
			return false
		}
	} else {
		o.channelRead <- &PacketWS{Status: Dataed, Socket: socket, Pkg:pkg}
		return true
	}
}

func (o *ServerWS) handlerServerWrite(socket uint32) bool {
	wsc, ok := o.connections[socket]
	if !ok {
		return false
	}

	select {
		case v, ok := <-wsc.channelWrite:
			if !ok {
				return false
			}

			err := websocket.JSON.Send(wsc.conn, *v.Pkg)
			if err != nil {
				fmt.Println("Server Write Packet Error ", err.Error(), wsc.socket)
				return false
			}
			return true
		default:
			return true
	}
}

func (o *ServerWS) handlerConnection(ws *websocket.Conn) {
	fmt.Println("New Connection Come ...")

	wsc := &WSClient{
		socket: WSAutoAmount(),
		conn: ws,
		channelStop: make(chan bool),
		channelWrite: make(chan *PacketWS, common.DefaultChannelLen),
	}

	o.connections[wsc.socket] = wsc

	o.channelRead <- &PacketWS{Status: Connected, Socket: wsc.socket}

	defer close(wsc.channelStop)
	defer delete(o.connections, wsc.socket)
	defer wsc.conn.Close()

	for o.handlerServerWrite(wsc.socket) && o.handlerServerRead(wsc.socket) {
		time.Sleep(time.Millisecond)
	}
}

func (o *ServerWS) handlerListen() {
	addr := fmt.Sprintf("%s:%d", o.ip, o.port)

	o.srv = &http.Server{
		Addr: addr,
		Handler: websocket.Handler(o.handlerConnection),
	}
	o.srv.SetKeepAlivesEnabled(true)

	defer close(o.channelStop)
	defer o.cleanServer()

	if err := o.srv.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func (o *ServerWS) Start(ip string, port uint) chan *PacketWS {
	o.ip = ip
	o.port = port
	o.connections = make(map[uint32]*WSClient)
	o.channelStop = make(chan bool)
	o.channelRead = make(chan *PacketWS, common.DefaultChannelLen)

	go o.handlerListen()

	return o.channelRead
}

func (o *ServerWS) Stop() {
	o.srv.Close()

	<-o.channelStop
}

func (o *ServerWS) cleanServer() {
	for _, v := range o.connections {
		v.conn.Close()
		close(v.channelWrite)

		<-v.channelStop
	}
}

func (o *ServerWS) CloseWebsocketClient(socket uint32) {
	wsc, ok := o.connections[socket]
	if ok {
		wsc.conn.Close()
		close(wsc.channelWrite)
		delete(o.connections, socket)
	}
}

func (o *ServerWS) SendWebsocketPacket(socket uint32, pkg *map[string]interface{}) {
	dc, ok := o.connections[socket]
	if ok {
		dc.channelWrite <- &PacketWS{Status: Dataed, Socket: socket, Pkg:pkg}
	}
}

func (o *ServerWS) GetClientAddr(socket uint32) string {
	dc, ok := o.connections[socket]
	if !ok {
		return ""
	}
	return dc.conn.RemoteAddr().String()
}