--组切换界面
SwitchViewBase = class("SwitchViewBase", SwallowViewBase)
local scheduler = cc.Director:getInstance():getScheduler()
switch_view_groups = switch_view_groups or {}
local schedulerID = nil
function SwitchViewBase.switchView(group, zorder, cls, ...)
	local options = {...}
	SwitchViewBase.switchViewFun(group,zorder,cls,nil,unpack(options))
end

function SwitchViewBase.switchViewFun(group, zorder, cls, fun, ...)
	local options = {...}
	myScheduler.performWithDelayGlobal(function()
		local v = switch_view_groups[group]
		SwitchViewBase.closeView(v)
		local scene = display.getRunningScene()
		local children = scene:getChildren()
	    for k,v in pairs(children) do
	        if v.super and v.super.__cname == "PopUpViewBase" then
	            v:close()
	        end
	    end
		local view = cls.new(unpack(options))
		global.g_currentViewName = view.__cname
		scene:addChild(view, zorder)
		view.__group__ = group
		switch_view_groups[group] = view
		if fun then
			fun()
		end
	end, 0)
end

function SwitchViewBase.closeView(view)
	if not view then
		return
	end
	view:close()
end

function SwitchViewBase:ctor(bSwallow, clr)
	SwitchViewBase.super.ctor(self,bSwallow,clr)
	self.maskLayer_:setPosition(0, 0)
end

function SwitchViewBase:close()
	switch_view_groups[self.__group__] = nil
	SwitchViewBase.super.close(self)
end

function SwitchViewBase:onEnter()
	-- printInfo("SwitchViewBase onEnter")
end

