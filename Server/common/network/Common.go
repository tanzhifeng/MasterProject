package network

import "net"

//Packet Define
const (
	LenHead = 2
	LenProt = 4
)

//Status
const (
	Connected = iota
	Losted
	Dataed
)

//Packet Data
type Packet struct {
	Status  int
	Socket  uint32
	Ptid    uint32
	Content []byte
}

type PacketWS struct {
	Status int
	Socket uint32
	Pkg *map[string]interface{}
}

//ReadFull ...
func ReadFull(conn *net.TCPConn, bytes []byte) (int, error) {
	lenNeed := len(bytes)
	lenRead := 0
	for lenRead < lenNeed {
		n, err := conn.Read(bytes[lenRead:])
		if err != nil {
			return -1, err
		}
		lenRead += n
	}
	return lenRead, nil
}

//WriteFull ...
func WriteFull(conn *net.TCPConn, bytes []byte) (int, error) {
	lenNeed := len(bytes)
	lenWrite := 0
	for lenWrite < lenNeed {
		n, err := conn.Write(bytes[lenWrite:])
		if err != nil {
			return -1, err
		}
		lenWrite += n
	}
	return lenWrite, nil
}

type HandlerC2S func(uint32, []byte)

type HandlerS2C func([]byte)

type HandlerWC2S func(uint32, map[string]interface{})