HallViewPanel = class("HallViewPanel",SwitchViewBase)
function HallViewPanel:ctor()
	HallViewPanel.super.ctor(self,true)
	print("进入大厅")
	local node = cc.CSLoader:createNode("csb/HallView/HallViewPanel.csb")
		:move(display.cx,display.cy)
		:addTo(self)
end


function HallViewPanel:onEnter()
	-- body
	print("HallViewPanel onEnter")
end

function HallViewPanel:onExit()
	-- body
	print("HallViewPanel onExit")
end