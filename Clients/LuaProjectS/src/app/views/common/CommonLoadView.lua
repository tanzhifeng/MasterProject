CommonLoadView = class("CommonLoadView", SwallowViewBase)
local instance = instance or nil
function CommonLoadView.getInstance()
	instance = instance or CommonLoadView.new()
	return instance
end

function CommonLoadView:destoryInstance()
	if instance then
		instance:release()
		instance = nil
	end
end

function CommonLoadView:ctor()
	CommonLoadView.super.ctor(self, true)
	self:retain()
	self.node = cc.CSLoader:createNode("csb/Common/CommonLoadPanel.csb")
		:addTo(self)
	self.inofText = self.node:getChildByName("Text_1")
	self.textCount = 1
	self.Text = {
	"网络不稳定连接已断开，正尝试重连",
	"网络不稳定连接已断开，正尝试重连.",
	"网络不稳定连接已断开，正尝试重连..",
	"网络不稳定连接已断开，正尝试重连...",
	}
end

function CommonLoadView:showCommonLoad()
	self.conectCount = 0
	self:closeCommonLoad()
	instance:setPosition(cc.p(display.cx, display.cy))
	local scene = display.getRunningScene()
	scene:addChild(instance, 888)
	self.noticeTick_ = myScheduler.scheduleGlobal(function() self:onTextUpdate() end, 0.5)
end
function CommonLoadView:onTextUpdate(txid)
		self.conectCount = self.conectCount + 1
	   	if self.conectCount%6 == 0 then --计时器三秒链接一次
	   		global.g_connect_count = global.g_connect_count +1
	   		if global.g_connect_count >= 6 or global.g_connect_5min_count > 8 then 
	   			local closeGame = CommonCloseGame:getInstance()
	   			closeGame:updataUI()
	   			closeGame:showCommonCloseGame()
	   			CommonLoadView:getInstance():closeCommonLoad()
		        return
	   		end
	   		
	   		if not netmng.g_gs_conn then
		   		global.g_reconnection = true --是否是重连标示
		   		netmng.g_gs_sock:close()
				netmng.g_gs_conn = false
				netmng.g_gs_init = false
				netmng.setGsNetAddress(netmng.GS_SERV_IP,netmng.GS_SERV_PORT)
			    netmng.gsInit()
			end
		end
	self.inofText:setString(self.Text[self.textCount%4+1])
	self.textCount = self.textCount + 1
end

function CommonLoadView:closeCommonLoad()
	if self.noticeTick_ then
		myScheduler.unscheduleGlobal(self.noticeTick_)
		self.noticeTick_ = nil
	end
	instance:removeSelf()
end