module( "gatesvr", package.seeall )

--------------- 解包 --------------------
function onNotifyRoomAdded(data)
	table.dump(data)

	global.g_mainPlayer:addRoom(data.item.roomid, data.item.roomname, data.item.roomkind, data.item.tablecount, data.item.chaircount)
	global.g_gameDispatcher:dispatcherEvent(EVENT_NOTIFY_ROOM_ADDED, data.item.roomid)
end

function onNotifyRoomRemoved(data)
	table.dump(data)

	global.g_mainPlayer:removeRoom(data.roomid)
	global.g_gameDispatcher:dispatcherEvent(EVENT_NOTIFY_ROOM_REMOVED, data.roomid)
end

function onNotifyTableAdded(data)
	table.dump(data)

	global.g_mainPlayer:addRoomTable(data.tableid, data.capacity)
	global.g_gameDispatcher:dispatcherEvent(EVENT_NOTIFY_ROOM_TABLE_ADDED, data.tableid, data.capacity)
end

function onNotifyTableRemoved(data)
	table.dump(data)

	global.g_mainPlayer:removeRoomTable(data.tableid)
	global.g_gameDispatcher:dispatcherEvent(EVENT_NOTIFY_ROOM_TABLE_REMOVED, data.tableid)
end

function onRoomUserEnter(data)
	table.dump(data)

	if data.code == codes.C_GT_SUCCESS then
		global.g_mainPlayer:setCurrentRoomId(data.roomid)

		global.g_mainPlayer:initRoomUser()
		for k, v in pairs(data.users) do
			global.g_mainPlayer:addRoomUser(v.userid, v.name, v.gold, v.tableid, v.chairid)
		end

		global.g_mainPlayer:initRoomTable()
		for k, v in pairs(data.tables) do
			global.g_mainPlayer:addRoomTable(v.tableid, v.capacity)
		end

		TextTipsUtils:showTips("进入房间成功 !")
		global.g_gameDispatcher:dispatcherEvent(EVENT_ROOM_USER_ENTER)
	else
		TextTipsUtils:showTips("进入房间失败 !")
	end
end

function onRoomUserLeave(data)
	table.dump(data)

	if data.code == codes.C_GT_SUCCESS then
		global.g_mainPlayer:setCurrentRoomId(nil)
		global.g_mainPlayer:initRoomUser()
		global.g_mainPlayer:initRoomTable()

		TextTipsUtils:showTips("退出房间成功 !")
		global.g_gameDispatcher:dispatcherEvent(EVENT_ROOM_USER_LEAVE)
	else
		TextTipsUtils:showTips("退出房间失败 !")
	end
end

function onNotifyRoomAppendUser(data)
	table.dump(data)

	TextTipsUtils:showTips("广播玩家进入房间 !")

	global.g_mainPlayer:addRoomUser(data.userid, data.name, data.gold, data.tableid, data.chairid)
	global.g_gameDispatcher:dispatcherEvent(EVENT_NOTIFY_ROOM_USER_ADDED, data.userid)
end

function onNotifyRoomRemoveUser(data)
	table.dump(data)

	TextTipsUtils:showTips("广播玩家退出房间 !")
	
	global.g_mainPlayer:removeRoomUser(data.userid)
	global.g_gameDispatcher:dispatcherEvent(EVENT_NOTIFY_ROOM_USER_REMOVED, data.userid)
end

function onRoomUserSitdown(data)
	table.dump(data)
end

function onRoomUserStandup(data)
	table.dump(data)
end

function onRoomNotifyUserSitdown(data)
	table.dump(data)
end

function onRoomNotifyUserStandup(data)
	table.dump(data)
end

--------------- 封包 --------------------
function sendEnterRoom(roomid)
	sendGtData(pid.GT_ROOM_USER_ENTER, {roomid = roomid})
end

function sendLeaveRoom()
	sendGtData(pid.GT_ROOM_USER_LEAVE)
end

function sendSitdown(tableid, chairid)
	sendGtData(pid.GT_ROOM_USER_SITDOWN, {tableid = tableid, chairid = chairid})
end

function sendStandup()
	sendGtData(pid.GT_ROOM_USER_STANDUP)
end