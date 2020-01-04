module( "gatesvr", package.seeall )

--------------- 解包 --------------------
function onLogin(data)
	

	dump(data)

	if data.code == codes.C_LOGIN_SUCCESS then
		-- beatsStart()

		-- adjustTime(data.time)

		-- global.g_mainPlayer:initPlayerData(data.userid, data.account, data.gold)
		-- global.g_mainPlayer:setCurrentRoomId(nil)
		-- global.g_mainPlayer:initRoomList()
		-- for k, v in pairs(data.rooms) do
		-- 	global.g_mainPlayer:addRoom(v.roomid, v.roomname, v.roomkind, v.tablecount, v.chaircount)
		-- end

		-- global.g_mainPlayer:initRoomTable()
		-- global.g_mainPlayer:initRoomUser()

		-- TextTipsUtils:showTips("登录成功 !")

		-- replaceScene(HallScene, TRANS_CONST.TRANS_SCALE)
		print("登录成功")
		SwitchViewBase.switchView(SWITCH_VIEW_GROUP_ONE,1,HallViewPanel) --跳到大厅
	else
		print("登录失败 ")
	end
end

function onRegister(data)
	
	dump(data)

	if data.code == codes.C_LOGIN_SUCCESS then
		print("注册成功 !")
		-- TextTipsUtils:showTips("注册成功 !")
	else
		print("注册失败 !")
		-- TextTipsUtils:showTips("注册失败 !")
	end
end

--------------- 封包 --------------------
function sendLogin(account, password)
	sendGtData(pid.GT_LOGIN, {account = account, password = password})
end

function sendRegister(account, password)
	sendGtData(pid.GT_REGISTER, {account = account, password = password})
end