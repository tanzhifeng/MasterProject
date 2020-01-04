package global

import (
	"../../common"
)

//获取所有在线房间
func (o *GlobalData) GetRooms() *map[uint64]*common.DataRoom {
	return &o.MapRoomData
}

func (o *GlobalData) IsRoomExist(roomid uint64) bool {
	_, ok := o.MapRoomData[roomid]

	return ok
}

func (o *GlobalData) AddRoomUser(roomid uint64, userid uint64) {
	u := o.GetUserData(userid)
	r := o.GetRoomByRoomId(roomid)

	u.RoomID = roomid
	r.RoomUsers[userid] = &common.DataSeat{
		TableID: 0,
		ChairID: 0,
	}
}

func (o *GlobalData) RemoveRoomUser(userid uint64) uint64 {
	u := o.GetUserData(userid)
	r := o.GetRoomByRoomId(u.RoomID)
	s := r.RoomUsers[userid]

	if s.TableID > 0 && s.ChairID > 0 {
		t := r.RoomTables[s.TableID]
		delete(t.Seats, s.ChairID)
	}

	roomid := u.RoomID

	delete(r.RoomUsers, userid)
	u.RoomID = 0

	return roomid
}

//获取在线房间
func (o *GlobalData) GetRoomByRoomId(roomid uint64) *common.DataRoom {
	return o.MapRoomData[roomid]
}

func (o *GlobalData) GetRoomBySocket(logicalsocket uint32) *common.DataRoom {
	roomid := o.MapRoomSocket[logicalsocket]

	return o.GetRoomByRoomId(roomid)
}