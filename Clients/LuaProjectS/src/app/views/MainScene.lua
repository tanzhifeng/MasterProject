local MainScene = class("MainScene", cc.load("mvc").ViewBase)
function MainScene:onCreate()
  
end

function MainScene:EventConnectSuccessed()
	print("EventConnectSuccessed")
end
function MainScene:EventConnectLosted()
	print("EventConnectLosted")
end
function MainScene:EventConnectFailed()
	print("EventConnectFailed")
end

function MainScene:onEnter()
	SwitchViewBase.switchView(SWITCH_VIEW_GROUP_ONE,1,LoginView)
	global.g_gameDispatcher:addEventListener(EVENT_CONNECT_SUCCESSED,self,self.EventConnectSuccessed)  
	global.g_gameDispatcher:addEventListener(EVENT_CONNECT_LOSTED,self,self.EventConnectLosted)  
	global.g_gameDispatcher:addEventListener(EVENT_CONNECT_FAILED,self,self.EventConnectFailed)  
end
function MainScene:onExit()
	global.g_gameDispatcher:removeEventListener(EVENT_CONNECT_SUCCESSED,self)
	global.g_gameDispatcher:removeEventListener(EVENT_CONNECT_LOSTED,self)
	global.g_gameDispatcher:removeEventListener(EVENT_CONNECT_FAILED,self)
end

return MainScene
