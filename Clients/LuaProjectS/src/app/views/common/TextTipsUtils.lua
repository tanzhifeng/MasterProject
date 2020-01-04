TextTipsUtils = class("TextTipsUtils")
local SHOW_TIPS_MAX = 3
local isDisplaying = false
local remainTips = {}
function TextTipsUtils:showTips(options,serverTips)
	if options == nil or options == "" then
		return
	end
	local param = nil
	if type(options) == "string" then
		if serverTips then 
			param = {text = "["..options.."]"}
		else
			param = {text = options}
		end
	else
		if serverTips then 
			options.text = "["..options.text.."]"
		end
		param = options
	end
	
	table.insert(remainTips, param)
	self:_displayTips(param)
end

function TextTipsUtils:_displayTips()
	if isDisplaying then
		return
	end
	isDisplaying = true
	if self.heartHandler then 
		myScheduler.unscheduleGlobal(self.heartHandler)
		self.heartHandler = nil
	end
	local options = table.remove(remainTips, 1)
	local nowScene = display.getRunningScene()
	local label = self:_createNormalTextLabel(options)
	label:align(display.CENTER, display.cx, display.cy )
	label:addTo(nowScene, 1000)
	local action = cc.Sequence:create(
		cc.DelayTime:create(options.time and options.time or 0),
		cc.Spawn:create(cc.FadeTo:create(2.0, 0),cc.MoveBy:create(3,cc.p(0, 160))),	
		cc.CallFunc:create(function() label:removeSelf() end))
	label:runAction(action)
	self:_nextStart()
end

function TextTipsUtils:_nextStart()
	self.heartHandler = myScheduler.scheduleGlobal(function()
        isDisplaying = false
		if #remainTips > 0 then
			self:_displayTips()
		end
    end, 0.5,false)
end

function TextTipsUtils:_createNormalTextLabel(options)
	local label  = cc.Label:createWithTTF(options.text, "fonts/gameFontALL.ttf", 40.0)
	label:enableOutline(cc.c4b(0,0,0,128),2)
	return label
end