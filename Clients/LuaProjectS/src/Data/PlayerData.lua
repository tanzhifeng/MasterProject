PlayerData = class("PlayerData")
function  PlayerData:ctor()
	if io.exists(cc.FileUtils:getInstance():getWritablePath().."gameData.dat") then
		-- 读取文件
		local logFile = io.open(cc.FileUtils:getInstance():getWritablePath().."gameData.dat", "r")
		local dataStr = logFile:read("*all")
		self.data_ = serializable.unserialize(dataStr)
		if type(self.data_) ~= "table" then 
			-- TextTipsUtils:showTips("初始化本地数据")
			self:initData()
		end 
		logFile:close()
		self:initSecond()
	else
		self:initData()
	end
end

-- 热更新后使用添加字段注意下bool字段
function PlayerData:initSecond()
	self.data_.clickValue = self.data_.clickValue and self.data_.clickValue or 1
	self:save()
end
function PlayerData:initData()
	--数据库第一次创建运行
	self.data_ = {
		refresh_token =  nil,		--setRefeshToken更改微信的刷新token
		is_show_Particle = true, 	--setParticleState粒子状态
		quick_enter_zoneId = "#888888",--默认期望进入的厅号
		music_open = true,--背景音乐是否打开
		sound_open = true,--音效是否打开
		music_level = 50,--背景音乐是否打开
		sound_level = 50,--音效是否打开
		bg_type = 1, --背景色
		sound_type = 1, --  1粤语 2国语
		guestId = "", --测试账号
		expect_club_id = "",--上一次进入的俱乐部房间
		noticeCount	= {},
		imageConfig = {},
		defalutIP = "127.0.0.1" ,
		sendHeat = 1,
		clickValue = 1,
		-----------------俱乐部排行表的排序方式
		sectionOrder = 1,
		sectionType = 1,
		dayOrder = 1,
		dayType = 1,
		monthOrder = 1,
		monthType = 1,
		------------------
		account = "123456",
		password = "123456",
	}
	self:save()
end


----------------把数据写入文件中
function PlayerData:save()
	-- body
	local open_file = io.open(cc.FileUtils:getInstance():getWritablePath().."gameData.dat","w")
	local data = serializable.serialize(self.data_)
	open_file:write(data)
	open_file:flush()
	open_file:close()
end


function PlayerData:setAccount(account)
	-- body
	self.data_.account = account
	self:save()
end

function PlayerData:getAccount()
	-- body
	return self.data_.account
end

function PlayerData:setPassword(password)
	-- body
	self.data_.password = password
	self:save()
end

function PlayerData:getPassword()
	-- body
	return self.data_.password
end


