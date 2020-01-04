package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"../../common"
)

var AmoutSocket uint32

type DataTCP struct {
	socket       uint32
	conn         *net.TCPConn
	channelStop chan bool
	channelWrite chan *Packet
}

type ServerTCP struct {
	ip          string
	port        uint
	listener    *net.TCPListener
	channelStop chan bool
	channelRead chan *Packet
	connections map[uint32]*DataTCP
}

func AutoAmount() uint32 {
	AmoutSocket++

	return AmoutSocket
}

func (o *ServerTCP) handlerServerRead(socket uint32) {
	fmt.Println("Start Server Read ...", socket)

	dc := o.connections[socket]

	defer close(dc.channelStop)
	defer delete(o.connections, socket)
	defer dc.conn.Close()

	for {
		bytesHead := make([]byte, LenHead)
		_, err := ReadFull(dc.conn, bytesHead)
		if err != nil {
			fmt.Println("Server Read Head Error ", err.Error(), dc.socket)
			o.channelRead <- &Packet{Status: Losted, Socket: socket}
			return
		}

		len := binary.BigEndian.Uint16(bytesHead)
		if len > 0 {
			bytesContent := make([]byte, len)
			_, err = ReadFull(dc.conn, bytesContent)
			if err != nil {
				fmt.Println("Server Content Read Error ", err.Error(), dc.socket)
				o.channelRead <- &Packet{Status: Losted, Socket: socket}
				return
			}

			ptid := binary.BigEndian.Uint32(bytesContent)

			o.channelRead <- &Packet{Status: Dataed, Socket: socket, Ptid: ptid, Content: bytesContent[LenProt:]}
		}
	}
}

func (o *ServerTCP) handlerServerWrite(socket uint32) {
	fmt.Println("Start Server Write ...", socket)

	dc := o.connections[socket]

	for {
		v, ok := <-dc.channelWrite
		if !ok {
			return
		}

		len := uint16(LenHead + LenProt + len(v.Content))
		bytesPacket := make([]byte, len)

		binary.BigEndian.PutUint16(bytesPacket, len-LenHead)
		binary.BigEndian.PutUint32(bytesPacket[LenHead:], v.Ptid)
		copy(bytesPacket[LenHead+LenProt:], v.Content)

		_, err := WriteFull(dc.conn, bytesPacket)
		if err != nil {
			fmt.Println("Server Write Packet Error ", err.Error())
			return
		}
	}
}

func (o *ServerTCP) handlerNewConnection(conn *net.TCPConn) {
	fmt.Println("New Connection Come ...")

	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(time.Second * 10)

	dc := &DataTCP{
		socket: AutoAmount(),
		conn: conn,
		channelStop: make(chan bool),
		channelWrite: make(chan *Packet, common.DefaultChannelLen),
	}

	o.connections[dc.socket] = dc

	o.channelRead <- &Packet{Status: Connected, Socket: dc.socket}

	go o.handlerServerRead(dc.socket)
	go o.handlerServerWrite(dc.socket)
}

func (o *ServerTCP) cleanServer() {
	for _, v := range o.connections {
		v.conn.Close()
		close(v.channelWrite)

		<-v.channelStop
	}
}

func (o *ServerTCP) handlerAcceptTCP() {
	defer close(o.channelStop)
	defer close(o.channelRead)
	defer o.cleanServer()

	for {
		conn, err := o.listener.AcceptTCP()
		if err != nil {
			fmt.Println("Server Accept Error ", err.Error())
			return
		}

		go o.handlerNewConnection(conn)
	}
}

func (o *ServerTCP) Start(ip string, port uint) bool {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("Server ResolveTCPAddr Error ", err.Error())
		return false
	}

	o.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("Server Listen Error ", err.Error())
		return false
	}

	o.ip = ip
	o.port = port
	o.connections = make(map[uint32]*DataTCP)
	o.channelStop = make(chan bool)
	o.channelRead = make(chan *Packet, common.DefaultChannelLen)

	go o.handlerAcceptTCP()

	return true
}

func (o *ServerTCP) Stop() {
	o.listener.Close()

	<-o.channelStop
}

func (o *ServerTCP) GetServerPacket() (*Packet, bool) {
	v, ok := <-o.channelRead

	return v, ok
}

//SendServerPacket Function
func (o *ServerTCP) SendServerPacket(socket uint32, ptid uint32, content []byte) {
	dc, ok := o.connections[socket]
	if ok {
		dc.channelWrite <- &Packet{Status: Dataed, Socket: socket, Ptid: ptid, Content: content}
	}
}

//CloseServerClient Function
func (o *ServerTCP) CloseServerClient(socket uint32) {
	dc, ok := o.connections[socket]
	if ok {
		dc.conn.Close()
		close(dc.channelWrite)
		delete(o.connections, socket)
	}
}

//GetClientAddr ...
func (o *ServerTCP) GetClientAddr(socket uint32) string {
	dc, ok := o.connections[socket]
	if !ok {
		return ""
	}
	return dc.conn.RemoteAddr().String()
}