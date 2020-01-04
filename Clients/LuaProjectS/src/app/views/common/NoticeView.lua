NoticeView = class("NoticeView", function() 
	local node = cc.Node:create()
	node:enableNodeEvents()
	return node 
end)

function NoticeView:ctor(braodWidth,SceneCode)
    
	self.braodWidth = braodWidth or 600
    self.label = cc.Label:createWithSystemFont("关于","Microsoft YaHei",25) 
         :setAnchorPoint(cc.p(0,0)) 
    local scrollViewLayer = cc.Layer:create():setPosition(cc.p(0,0)) 
    local scrollView1 = cc.ScrollView:create()

    if nil ~= scrollView1 then 
        scrollView1:setViewSize(cc.size(self.braodWidth, 100)) 
        scrollView1:setDirection(cc.SCROLLVIEW_DIRECTION_NONE ) 
        scrollView1:setClippingToBounds(true) 
        scrollView1:setBounceable(true)
        scrollView1:setTouchEnabled(false) 
    end 
    scrollView1:addChild( self.label) 
    scrollView1:setPositionX(-25)
    self:addChild(scrollView1) 

    if nil ~= scrollViewLayer_ then 
        scrollView1:setContainer(scrollViewLayer) 
        scrollView1:updateInset() 
    end 
end

function NoticeView:startNotice()
	self:stopNotice()
	self:onNoticeUpdate()
    self.noticeTick_ = myScheduler.scheduleGlobal(function() self:onNoticeUpdate() end, 1)
end


function NoticeView:onNoticeUpdate()
	if self.playing_ then
		return
	end
	local notice = global.g_playerData.noticeData:popNextNotice()
	if notice  then
        local noticeData = global.g_playerData.noticeData:getNotice(notice.notice_id)
        local nowTime = os.time()
        if nowTime < noticeData.end_time then
        else
            self:hide()
            return
        end
        global.g_playerData:addNoticesCount(notice.notice_id)
        self.playing_ = true
        
		self.label:setString(noticeData.content)

        local labelWidth =  self.label:getContentSize().width 
        local time = (labelWidth+self.braodWidth)/100 -- 这里可以根据label多长动态计算时间 
   
        self.label:setPosition(cc.p(self.braodWidth, 0)) 

        local leftAction = cc.MoveBy:create(time,cc.p(-labelWidth-self.braodWidth ,0))
        local delay = cc.DelayTime:create(3)
        local callbackAction = cc.CallFunc:create(function(sender) 
            self.playing_ = nil
        end)
        local seqAction = cc.Sequence:create(leftAction,delay,callbackAction)
        self.label:runAction(seqAction) 
		self:show()
	else
		self:hide()
	end
end

function NoticeView:stopNotice()
	if self.noticeTick_ then
        myScheduler.unscheduleGlobal(self.noticeTick_)
		self.noticeTick_ = nil
	end
end

function NoticeView:onEnter()
	self:startNotice()
end

function NoticeView:onExit()
	self:stopNotice()
end