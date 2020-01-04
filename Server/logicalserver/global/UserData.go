package global

import (
	"../../common"
)

func (o *GlobalData) GetLogins() *map[uint32]map[uint32]uint64 {
	return &o.MapLoginSocket
}

func (o *GlobalData) GetLoginSockets(logicalsocket uint32) *map[uint32]uint64 {
	v := o.MapLoginSocket[logicalsocket]
	return &v
}

//获取所有在线玩家
func (o *GlobalData) GetUsers() *map[uint64]*common.DataUser {
	return &o.MapUserData
}

//新增在线玩家
func (o *GlobalData) AddUser(logicalsocket uint32, loginsocket uint32, address string, userid uint64, account string, password string, gold uint64) {
	sockets, ok := o.MapLoginSocket[logicalsocket]
	if !ok {
		return
	}

	sockets[loginsocket] = userid

	o.MapUserData[userid] = &common.DataUser{
		LogicalSocket: logicalsocket,
		LoginSocket: loginsocket,
		Address: address,
		UserID: userid,
		Account: account,
		Password: password,
		Gold: gold,
		Status: common.USER_FREE,
		RoomID: 0,
	}
}

//通过玩家socket获取玩家id
func (o *GlobalData) GetUserIdBySocket(logicalsocket uint32, loginsocket uint32) uint64 {
	sockets, ok := o.MapLoginSocket[logicalsocket]
	if !ok {
		return 0
	}

	return sockets[loginsocket]
}

//获取玩家数据
func (o *GlobalData) GetUserData(userid uint64) *common.DataUser {
	return o.MapUserData[userid]
}

func (o *GlobalData) IsUserOnlineByUserId(userid uint64) bool {
	_, ok := o.MapUserData[userid]

	return ok
}

func (o *GlobalData) IsUserOnlineBySocket(logicalsocket uint32, loginsocket uint32) bool {
	userid := o.GetUserIdBySocket(logicalsocket, loginsocket)

	return o.IsUserOnlineByUserId(userid)
}

func (o *GlobalData) IsRoomUserByUserId(userid uint64) bool {
	v := o.GetUserData(userid)

	return v.RoomID > 0
}

func (o *GlobalData) IsRoomUserBySocket(logicalsocket uint32, loginsocket uint32) bool {
	userid := o.GetUserIdBySocket(logicalsocket, loginsocket)

	return o.IsRoomUserByUserId(userid)
}

//删除已登录帐号
func (o *GlobalData) RemoveUserByUserId(userid uint64) {
	userdata, ok := o.MapUserData[userid]
	if ok {
		delete(o.MapUserData, userid)

		sockets, ok := o.MapLoginSocket[userdata.LogicalSocket]
		if !ok {
			return
		}
		delete(sockets, userdata.LoginSocket)
	}
}