-- 模态界面,吞噬触摸事件
SwallowViewBase = class("SwallowViewBase", EventNode)

function SwallowViewBase.showSwallowView(cls, ...)
	local scene = display.getRunningScene()
	local popup = cls.new(...)
		:move(display.cx, display.cy)
		:addTo(scene, 999) 
	return popup
end

function SwallowViewBase.showSwallowPanel(cls, x, y, ...)
	local scene = display.getRunningScene()
	local popup = cls.new(...)
		:move(x, y)
		:addTo(scene, 999)
	return popup
end


function SwallowViewBase:close()
	self:removeSelf()
end

function SwallowViewBase:ctor(bSwallow, clr)
	SwallowViewBase.super.ctor(self, bSwallow, clr)
	self.maskLayer_ = cc.LayerColor:create(clr or cc.c4b(0, 0, 0, 0))
	self.maskLayer_:setPosition(-display.width/2, -display.height/2)
    self:addChild(self.maskLayer_,-2)
    
    if bSwallow then
		self.maskLayer_:registerScriptTouchHandler(function(evt, x, y)
	        if evt == "began" then
	        	return self:onTouchBegan(x, y)
	        elseif evt == "moved" then
	        	self:onTouchMoved(x, y)
	        elseif evt == "ended" then
	        	self:onTouchEnded(x, y)
	        elseif evt == "canceled" then
	        	self:onTouchCanceled(x, y)
	        end
	    end)
	    self.maskLayer_:setTouchEnabled(true)
	    self.maskLayer_:setSwallowsTouches(true)
    end
end

function SwallowViewBase:onTouchBegan(x, y)
	--子类重写
	return true
end

function SwallowViewBase:onTouchMoved(x, y)
	--子类重写
end

function SwallowViewBase:onTouchEnded(x, y)
	--子类重写
end

function SwallowViewBase:onTouchCanceled(x, y)
	--子类重写
end

function SwallowViewBase:onEnter()
	-- printUserInfo("onEnter SwitchView")
end
