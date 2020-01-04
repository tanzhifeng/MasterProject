--弹出界面
PopUpViewBase = class("PopUpViewBase", SwallowViewBase)

function PopUpViewBase.showPopUpView(cls, ...)
	local scene = display.getRunningScene()
	local popup = cls.new(...)
		:move(display.cx, display.cy)
		popup:addTo(scene, 999)
	popup.nodeUI_:setScale(0)

	local action = cc.Sequence:create(
		cc.ScaleTo:create(0.1, 1),
		cc.CallFunc:create(handler(popup,popup.onPopUpComplete)))
	popup.nodeUI_:runAction(action)
	return popup
end

function PopUpViewBase.showPopUpViewWithBlur(cls, ...)
	local scene = display.getRunningScene()
	local popup = cls.new(...)
		:move(display.cx, display.cy)
		:addTo(scene, 999)
		getScreenShotBlur(10, 7, 660, function(blurShot)
			popup:addChild(blurShot, -1)
			local action = cc.Sequence:create({
					cc.ScaleTo:create(0.1, 1),
					cc.CallFunc:create(handler(popup, popup.onPopUpComplete))
				})
			popup.nodeUI_:runAction(action)
		end)
	return popup
end


function PopUpViewBase.closePopUpView(view)
	local action = cc.Sequence:create(
		-- cc.ScaleTo:create(0.1, 0),
		cc.CallFunc:create(function()
		view:onPopBackComplete()
		view:removeSelf()
				end))
	view.nodeUI_:runAction(action)
end

function PopUpViewBase:onPopUpComplete()
	--弹出效果结束后调用
	-- printUserInfo("打开界面完成")
end

function PopUpViewBase:onPopBackComplete()
	--关闭效果结束后调用
	-- printUserInfo("关闭界面完成")
end

function PopUpViewBase:ctor(bSwallow, clr)
	PopUpViewBase.super.ctor(self, bSwallow, clr)

	self.nodeUI_ = display.newNode()
		:addTo(self, 2)
end

function PopUpViewBase:close()
	--子类可继承,但要调用基类
	PopUpViewBase.closePopUpView(self)

end


function PopUpViewBase:onEnter()
	
end
function PopUpViewBase:onExit( )
	
end
