CommonCloseGame = class("CommonCloseGame", SwallowViewBase)

local instance = instance or nil
function CommonCloseGame.getInstance()
	instance = instance or CommonCloseGame.new()
	return instance
end

function CommonCloseGame:showCommonCloseGame(confirmFunc)
	instance:removeSelf()
	if confirmFunc then
		self.confirmFunc_ = confirmFunc
	end
	instance:setPosition(cc.p(display.cx, display.cy))
	local scene = display.getRunningScene()
	scene:addChild(instance, 1001)
end

function CommonCloseGame:ctor()
	CommonCloseGame.super.ctor(self, true, cc.c4b(0, 0, 0, 128))

	local node = cc.CSLoader:createNode("csb/Common/CommonConfirmPanel.csb")
		:addTo(self)
	self:retain()
	self.labelContent = node:getChildByName("Text_1")
	self.labelContent:setTextHorizontalAlignment(1) --对齐方式

	local btnConfirm = node:getChildByName("Button_1")
	btnConfirm:addTouchEventListener(makeClickHandler(self, self.onConfirm))
	local btnCancel  = node:getChildByName("Button_2")
	btnCancel:addTouchEventListener(makeClickHandler(self, self.onCancel))
	self.confirmFunc_ = function() os.exit() end





	
	btnCancel:setVisible(false)
	btnConfirm:setVisible(false)

	btnConfirm:setVisible(true)
	btnConfirm:setPositionX(0)
end

function CommonCloseGame:updataUI(text)
	if text then 
		self.labelContent:setString(text)
	else
		self.labelContent:setString("网络不稳定，请检查网络!")
	end
	local labelVr = self.labelContent:getVirtualRenderer()
	labelVr:setDimensions(420, 0)
	local lsize = labelVr:getContentSize()
	self.labelContent:ignoreContentAdaptWithSize(false)
	self.labelContent:setTextAreaSize(lsize)
end

function CommonCloseGame:onCancel()
	if self.cancelFunc_ then
		self.cancelFunc_()
	end
	self:close()
end

function CommonCloseGame:onConfirm()
	if self.confirmFunc_ then
		self.confirmFunc_()
	end
	self:close()
end

function CommonCloseGame:onDelect()
	if self.delectFunc_ then
		self.delectFunc_()
	end
	self:close()
end


function CommonCloseGame:onEnter()
	global.g_gaming = 0
	netmng.g_gt_init = false
	gamesvr.stopHeart()
	netmng.g_gs_sock:close()	
 	global.g_gameDispatcher:dispatchEvent(EVENT_CLEAN_HIDE_CREATEROOM_TEXTFILE,false)	
end


function CommonCloseGame:onExit()
	global.g_gameDispatcher:dispatchEvent(EVENT_CLEAN_HIDE_CREATEROOM_TEXTFILE,true)
end