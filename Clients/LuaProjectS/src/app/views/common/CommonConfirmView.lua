CommonConfirmView = class("CommonConfirmView", PopUpViewBase)

function CommonConfirmView.showCommonConfirmView(param)
	return PopUpViewBase.showPopUpView(CommonConfirmView, param)
end

function CommonConfirmView:ctor(param)
	CommonConfirmView.super.ctor(self, true, cc.c4b(0, 0, 0, 170))
	local node = cc.CSLoader:createNode("csb/Common/CommonConfirmPanel.csb")
		:addTo(self.nodeUI_)
		
	param.fontSize = param.fontSize or 38
	local labelContent = node:getChildByName("Text_1")
	labelContent:setTextHorizontalAlignment(param.align or cc.TEXT_ALIGNMENT_CENTER) --对齐方式
	labelContent:setString(param.text)
	labelContent:setFontSize(param.fontSize)
	labelContent:setTextAreaSize(cc.size(550,0))
	local btnConfirm = node:getChildByName("Button_1")
	btnConfirm:addTouchEventListener(makeClickHandler(self, self.onConfirm))
	local btnCancel  = node:getChildByName("Button_2")
	btnCancel:addTouchEventListener(makeClickHandler(self, self.onCancel))
	local tipsText = node:getChildByName("Text_3")
	if param.tipsText then 
		tipsText:setString(param.tipsText)
	else
		tipsText:setString("")
	end
	self.cancelFunc_ = param.cancel
	self.confirmFunc_ = param.confirm
	param.style = param.style or COMMON_CONFIRM_YN
	btnCancel:setVisible(false)
	btnConfirm:setVisible(false)
	if param.style == COMMON_CONFIRM_YN or param.style == COMMON_CONFIRM_NY then
		btnCancel:setVisible(true)
		btnConfirm:setVisible(true)
		btnCancel:setPositionX(-150.0)
		btnConfirm:setPositionX(150.0)
	elseif param.style == COMMON_CONFIRM_N then
		btnCancel:setVisible(true)
		btnCancel:setPositionX(0)
	else
		btnConfirm:setVisible(true)
		btnConfirm:setPositionX(0)
	end
end

function CommonConfirmView:onCancel()
	if self.cancelFunc_ then
		self.cancelFunc_()
	end
	self:close()
end

function CommonConfirmView:onConfirm()
	if self.confirmFunc_ then
		self.confirmFunc_()
	end
	self:close()
end
