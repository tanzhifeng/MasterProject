LoginView = class("LoginView",SwitchViewBase)
function LoginView:ctor()
	LoginView.super.ctor(self,true,cc.c4b(0,0,0,128))
	local node = cc.CSLoader:createNode("csb/LoginViewPanel/LoginViewPanel.csb")
        :move(display.cx,display.cy)
        :addTo(self)
    self.nameText = node:getChildByName("AccountTextField")
    self.passText = node:getChildByName("PasswordTextField")

    local loginBut = node:getChildByName("LoginButton")
    loginBut:addTouchEventListener(makeClickHandler(self,self.loginFun))

    local regsBut = node:getChildByName("RegistButton")
    regsBut:addTouchEventListener(makeClickHandler(self,self.regsFun))
end
-- c2s_login_user
function LoginView:loginFun()
    gatesvr.startGate()
	-- gatesvr.sendGtData(pid.GT_LOGIN,{account = self.nameText:getString(),password = self.passText:getString()})
	-- print(self.nameText:getString())
	-- print(self.passText:getString())
    global.g_playerData:setAccount(self.nameText:getString())
    global.g_playerData:setPassword(self.passText:getString())
end

function LoginView:regsFun()
    -- gatesvr.startGate()
	gatesvr.sendGtData(gatesvr.pid.LOGIN_REGISTER,{account = self.nameText:getString(),password = self.passText:getString()})
    print(self.nameText:getString())
    print(self.passText:getString())
end