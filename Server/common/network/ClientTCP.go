package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"../../common"
)

type ClientTCP struct {
	ip string
	port uint
	conn *net.TCPConn
	channelStop chan bool
	channelRead chan *Packet
	channelWrite chan *Packet
}

func (o *ClientTCP) handlerRead() {
	defer close(o.channelStop)
	defer o.clean()

	for {
		bytesHead := make([]byte, LenHead)
		_, err := ReadFull(o.conn, bytesHead)
		if err != nil {
			fmt.Println("Client Read Head Error ", err.Error())
			o.channelRead <- &Packet{Status: Losted}
			return
		}

		len := binary.BigEndian.Uint16(bytesHead)
		if len > 0 {
			bytesContent := make([]byte, len)
			_, err = ReadFull(o.conn, bytesContent)
			if err != nil {
				fmt.Println("Client Read Content Error ", err.Error())
				o.channelRead <- &Packet{Status: Losted}
				return
			}

			ptid := binary.BigEndian.Uint32(bytesContent)

			o.channelRead <- &Packet{Status: Dataed, Socket: 0, Ptid: ptid, Content: bytesContent[LenProt:]}
		}
	}
}

func (o *ClientTCP) handlerWrite() {
	for {
		v, ok := <-o.channelWrite
		if !ok {
			return
		}

		len := uint16(LenHead + LenProt + len(v.Content))
		bytesPacket := make([]byte, len)

		binary.BigEndian.PutUint16(bytesPacket, len-LenHead)
		binary.BigEndian.PutUint32(bytesPacket[LenHead:], v.Ptid)
		copy(bytesPacket[LenHead+LenProt:], v.Content)

		_, err := WriteFull(o.conn, bytesPacket)
		if err != nil {
			fmt.Println("Client Write Packet Error ", err.Error())
			return
		}
	}
}

func (o *ClientTCP) Start(ip string, port uint) bool {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("Client ResolveTCPAddr Error ", err.Error())
		return false
	}

	o.conn, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("Client DialTCP Error ", err.Error())
		return false
	}

	o.ip = ip
	o.port = port
	o.channelStop = make(chan bool)
	o.channelRead = make(chan *Packet, common.DefaultChannelLen)
	o.channelWrite = make(chan *Packet, common.DefaultChannelLen)

	o.conn.SetKeepAlive(true)
	o.conn.SetKeepAlivePeriod(time.Second * 10)

	o.channelRead <- &Packet{Status: Connected}

	go o.handlerRead()
	go o.handlerWrite()

	return true
}

func (o *ClientTCP) Stop() {
	o.conn.Close()

	<-o.channelStop
}

func (o *ClientTCP) reconnect() bool {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", o.ip, o.port))
	if err != nil {
		fmt.Println("Client ResolveTCPAddr Error ", err.Error())
		return false
	}

	o.conn, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("Client DialTCP Error ", err.Error())
		return false
	}

	o.conn.SetKeepAlive(true)
	o.conn.SetKeepAlivePeriod(time.Second * 10)

	o.channelRead <- &Packet{Status: Connected}

	go o.handlerRead()
	go o.handlerWrite()

	return true
}

func (o *ClientTCP) clean() {
	o.conn.Close()
	close(o.channelRead)
	close(o.channelWrite)
}

func (o *ClientTCP) GetLocalAddr() string {
	return o.conn.LocalAddr().String()
}

func (o *ClientTCP) GetRemoteAddr() string {
	return o.conn.RemoteAddr().String()
}

func (o *ClientTCP) GetClientPacket() (*Packet, bool) {
	v, ok := <-o.channelRead

	return v, ok
}

func (o *ClientTCP) SendClientPacket(ptid uint32, content []byte) {
	o.channelWrite <- &Packet{Status: Dataed, Socket: 0, Ptid: ptid, Content: content}
}
