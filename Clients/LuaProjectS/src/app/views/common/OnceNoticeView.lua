-- 一次性通知
OnceNoticeView = class("OnceNoticeView",EventNode)

function OnceNoticeView:ctor(braodWidth)
    OnceNoticeView.super.ctor(self)
	self.braodWidth = braodWidth or 600
    self.noticeList = {}

    local node = cc.CSLoader:createNode("csb/Common/OnceNoticeView.csb")
        :addTo(self)

    local bg_sprite = node:getChildByName("Image_1")
    bg_sprite:setContentSize(self.braodWidth+80, bg_sprite:getContentSize().height)    
    local panelLayer = node:getChildByName("Panel_1")
    panelLayer:setContentSize(cc.size(self.braodWidth, bg_sprite:getContentSize().height))
    self.content = {}

    self.content[1] = panelLayer:getChildByName("Text_1")
    self.content[1]:setPositionY(0.5*bg_sprite:getContentSize().height)

    self.content[2] = panelLayer:getChildByName("ListView_2")

    self.textList = {}
    self.textList[1] = self.content[2]:getChildByName("UserName_1")
    self.textList[2] = self.content[2]:getChildByName("Text_9")
    self.textList[3] = self.content[2]:getChildByName("UserName_2")
    self.textList[4] = self.content[2]:getChildByName("Text_11")
    self.textList[5] = self.content[2]:getChildByName("Distance")
    self:hide() 
end

function OnceNoticeView:addOnceNotice(notice)
    table.insert(self.noticeList,notice)
    if not self.noticeTick_ then
        self:onNoticeUpdate()
        self.noticeTick_ = myScheduler.scheduleGlobal(function() self:onNoticeUpdate() end, 0.5)
    end
end

function OnceNoticeView:onNoticeUpdate()
	if self.playing_ then
		return
	end
	local notice = table.remove(self.noticeList, 1)
	if notice then
        self.playing_ = true
        for i,v in ipairs(self.content) do
            v:hide()
        end
        local labelWidth = 0
        if type(notice) == "string" then -- 一般性通知
            self.content[1]:show()
            self.content[1]:setString(notice)
            labelWidth = self.content[1]:getContentSize().width
            self.content[1]:setPosition(cc.p(self.braodWidth,self.content[1]:getPositionY()))
            self:objRunAction(self.content[1],labelWidth)
        elseif notice.type == 1 then
            self.content[1]:show()
            self.content[1]:setString(notice.content)
            labelWidth = self.content[1]:getContentSize().width
            self.content[1]:setPosition(cc.p(self.braodWidth,self.content[1]:getPositionY()))
            self:objRunAction(self.content[1],labelWidth)
        elseif notice.type == 2 then  --距离警告通知
            self.content[2]:show()
            self.textList[1]:setString(notice.nickName1)
            self.textList[2]:setString("与")
            self.textList[3]:setString(notice.nickName2)
            self.textList[4]:setString("的距离仅有")
            self.textList[5]:setString(notice.distance.."米")
            for i,v in ipairs(self.textList) do
                labelWidth = labelWidth+v:getContentSize().width
            end
            self.content[2]:setPosition(cc.p(self.braodWidth,self.content[2]:getPositionY())) 
            self:objRunAction(self.content[2],labelWidth)
        else

        -- elseif notice.type = 3 then 
        --     self.content[2]:show()
        --     self.textList[1]:setString("")
        --     self.textList[2]:setString("")
        --     self.textList[3]:setString("")
        --     self.textList[4]:setString("")
        --     self.textList[5]:setString("")
        --     for i,v in ipairs(self.textList) do
        --         labelWidth = labelWidth+v:getContentSize().width
        --     end
        --     self.content[2]:setPosition(cc.p(self.braodWidth,self.content[2]:getPositionY())) 
        --     self:objRunAction(self.content[2],labelWidth)
        end
		self:show()
	else
        self:stopNotice()
		self:hide()
	end
end

function OnceNoticeView:objRunAction(senderObj,labelWidth)
    local time = (labelWidth+self.braodWidth)/100 -- 这里可以根据label多长动态计算时间 
    local leftAction = cc.MoveBy:create(time,cc.p(-labelWidth-self.braodWidth ,0))
    local callbackAction = cc.CallFunc:create(function(sender) 
        self.playing_ = nil
    end)
    local seqAction = cc.Sequence:create(leftAction,delay,callbackAction)
    senderObj:runAction(seqAction)
end

function OnceNoticeView:stopNotice()
	if self.noticeTick_ then
        myScheduler.unscheduleGlobal(self.noticeTick_)
		self.noticeTick_ = nil
	end
end

function OnceNoticeView:onExit()
	self:stopNotice()
end