package global

import (
	"../../common"
	"sync"
)

var globaldata *GlobalData
var globaldataOnce sync.Once

type GlobalData struct {
	MapNodeData map[uint32]*common.DataNode //logicalsocket -- datanode

	MapLoginSocket map[uint32]map[uint32]uint64 //logicalsocket & loginsocket -- userid
	MapUserData map[uint64]*common.DataUser //userid -- userdata

	MapRoomSocket map[uint32]uint64 //logicalsocket -- roomid
	MapRoomData map[uint64]*common.DataRoom //roomid -- roomdata
}

func GetGlobalData() *GlobalData {
	globaldataOnce.Do(func() {
		globaldata = &GlobalData{}
		globaldata.init()
	})
	return globaldata
}

func (o *GlobalData) init() {
	o.MapNodeData = make(map[uint32]*common.DataNode)

	o.MapLoginSocket = make(map[uint32]map[uint32]uint64)
	o.MapUserData = make(map[uint64]*common.DataUser)

	o.MapRoomSocket = make(map[uint32]uint64)
	o.MapRoomData = make(map[uint64]*common.DataRoom)
}

