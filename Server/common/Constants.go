package common

//DefaultChannelLen ...
const DefaultChannelLen = 100

//Variant ...
const (
	BeatsIntervalServer = 15 //sec
)

//user status
const (
	USER_FREE = iota
	USER_GAME
	USER_TRUST
)

//Socket Type
const (
	SOCKET_UNKNOW = iota
	SOCKET_USER
)

type DataSocket struct {
	Socket uint32
	SocketType int32
	Address string
	Extra string
	BeatsLast int64
}

//SocketUser ...
type SocketUser struct {
	Socket uint32
	UserID uint64
	Account string
	Gold uint64
}

//SocketRoom ...
type SocketRoom struct {
	Socket uint32
	RoomID string
}

//位置数据
type DataSeat struct {
	TableID uint32
	ChairID uint32
}

//桌子数据
type DataTable struct {
	TableID uint32
	Capacity uint32
	Seats map[uint32]uint64
}

//房间数据
type DataRoom struct {
	LogicalSocket uint32
	RoomID uint64
	RoomName string
	RoomKind uint32
	RoomUsers map[uint64]*DataSeat
	RoomTables map[uint32]*DataTable
}

type RoomUserStruct struct {
	UserID uint64
	Name string
	Gold uint64
	TableID uint32
	ChairID uint32
}

type RoomTableStruct struct {
	TableID uint32
	Capacity uint32
}

//玩家数据
type DataUser struct {
	LogicalSocket uint32
	LoginSocket uint32
	Address string
	UserID uint64
	UserName string
	Account string
	Password string
	Gold uint64
	Status uint8
	RoomID uint64
}

//客户端数据记录
type DataNode struct {
	Socket uint32
	Address string
	NodeType uint32
}

//Component Identify
const (
	Scheduler = iota
	Logical
	LogicalClient
	LoginServer
	RoomClient
	RoomServer
	DatabaseClient
	LogicalServer
	MySQLClient
	DatabaseServer
)

type IComponent interface {
	Stop()
}