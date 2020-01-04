package global

import (
	"../../common"
	"../../common/prot/pt_com"
)

func (o *GlobalData) AddNodeClient(logicalsocket uint32, address string) {
	o.MapNodeData[logicalsocket] = &common.DataNode{
		Socket: logicalsocket,
		Address: address,
		NodeType: uint32(pt_com.NodeType_NODE_UNKNOW),
	}
}

func (o *GlobalData) GetNodeClient(logicalsocket uint32) *common.DataNode {
	return o.MapNodeData[logicalsocket]
}

func (o *GlobalData) RemoveNodeClient(logicalsocket uint32) {
	delete(o.MapNodeData, logicalsocket)
}

//新增登录节点
func (o *GlobalData) AddLoginServerNode(logicalsocket uint32) {
	v, ok := o.MapNodeData[logicalsocket]
	if !ok {
		return
	}
	v.NodeType = uint32(pt_com.NodeType_NODE_LOGIN_SERVER)
	o.MapLoginSocket[logicalsocket] = make(map[uint32]uint64)
}

//删除登陆节点
func (o *GlobalData) RemoveLoginServerNode(logicalsocket uint32) {
	delete(o.MapLoginSocket, logicalsocket)
	delete(o.MapNodeData, logicalsocket)
}

//新增房间节点
func (o *GlobalData) AddRoomServerNode(logicalsocket uint32, roomid uint64, roomname string, roomkind uint32) {
	v, ok := o.MapNodeData[logicalsocket]
	if !ok {
		return
	}
	v.NodeType = uint32(pt_com.NodeType_NODE_ROOM_SERVER)

	o.MapRoomSocket[logicalsocket] = roomid
	o.MapRoomData[roomid] = &common.DataRoom{
		LogicalSocket: logicalsocket,
		RoomID: roomid,
		RoomName: roomname,
		RoomKind: roomkind,
		RoomUsers: make(map[uint64]*common.DataSeat),
		RoomTables: make(map[uint32]*common.DataTable),
	}
}

//删除房间节点
func (o *GlobalData) RemoveRoomServerNode(logicalsocket uint32) {
	roomid, ok := o.MapRoomSocket[logicalsocket]
	if ok {
		delete(o.MapRoomData, roomid)
		delete(o.MapRoomSocket, logicalsocket)
		delete(o.MapNodeData, logicalsocket)
	}
}