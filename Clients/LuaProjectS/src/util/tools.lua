-- start --
--------------------------------
-- 计数器 加一减一处理
-- @function getIncOrDecNumber
-- end --
function getIncOrDecNumber()
	local t = {
		num = 0,
		inc = function(self)
			self.num = self.num + 1
			return self.num
		end,

		dec = function(self)
			self.num = self.num - 1
			return self.num
		end,
	}
	return t
end
-- start --
--------------------------------
-- 点击事件的响应方式
-- @function makeClickHandler
-- @param obj 点击的函数实体 func 点击的回调函数 sender 点击的实体
-- end --
function makeClickHandler(obj, func)
	return function(sender, event) 
        if event == 2 then
            func(obj, sender, event)
        end
    end
end
--服务器位置转客户端位置
--@targetSide 服务器目标位置
--@selfSeatIndex 服务器自己位置
--@MaxPlayer 最大玩家人数
function getClientSideID(targetSide,selfSeatIndex,MaxPlayer)
    -- return (targetSide-selfSeatIndex+MaxPlayer) % MaxPlayer + 1
    return (MaxPlayer - selfSeatIndex + targetSide) % MaxPlayer + 1
end


-- 点击事件的响应方式
-- @function makeChekcHandler
-- @param obj 点击的函数实体 func 点击的回调函数 sender 点击的实体
-- end --
function makeChekcHandler(obj, func)
    return function(sender, event) 
        if event == 1 or event == 0 then 
        elseif event == 2 then 
            func(obj, sender, event)
        elseif 3 == event then 
            sender:setSelected(not sender:isSelected())
        end 
    end
end


-- start --
--------------------------------
-- 以调整数字的缩小格式化
-- @function moneyFormat
-- @param num 数字
-- @
-- end --
local MONEY_FORMATS = {
    {100000000, "亿"},
    {10000000, "千万"},
    {1000000, "百万"},
    {100000, "十万"},
    {10000, "万"},
    {1, ""},
}
function moneyFormat(num)
    if type(num) == "string" then
        return num
    end
    if num < 0 then
        return "0"
    end
    local words = tostring(num)
    for i = 1, #MONEY_FORMATS do
        local data = MONEY_FORMATS[i]
        if num >= data[1] then
            local p = math.floor(num*100/data[1])/100
            words = tostring(p) .. data[2]
            break
        end
    end
    return words
end


function abbreviationFormat(number,unitStr)
    if type(number) == "string" then
        return number
    end
    local words = tostring(number)
    if number < 0 then
        for i = 1, #MONEY_FORMATS do
            local data = MONEY_FORMATS[i]
            if math.abs(number) >= data[1] then
                local p = math.floor(math.abs(number)*100/data[1])/100
                words = tostring(p) .. data[2]
                break
            end
        end
        words = "-"..words
    else
        for i = 1, #MONEY_FORMATS do
            local data = MONEY_FORMATS[i]
            if number >= data[1] then
                local p = math.floor(number*100/data[1])/100
                words = tostring(p) .. data[2]
                break
            end
        end
    end
    if unitStr then 
        words = words..unitStr
    end
    return words
end

--------------------------------
-- 以单选按钮的事件调整
-- @function CheckBoxChageEvnet
-- @param sender 按下的按钮
-- @param eventType 按下的类型
-- switch_but 按钮数组
-- switch_but_text 按钮文字数组
-- switch_node  翻页的数组
-- end --
function CheckBoxChageEvnet(sender,eventType,switch_but,switch_node,switch_but_text)
    if eventType == 2 then
        if switch_node then 
            if sender:isSelected() then
                for k,v in pairs(switch_node) do
                    if sender:getTag() == k then
                        v:setVisible(true)
                    else
                        v:setVisible(false)
                    end
                end
            end
        end

        if switch_but_text then
            for k,v in pairs(switch_but_text ) do
                if sender:getTag() == k then
                    v:setSelected(true)
                else
                    v:setSelected(false)
                end
            end
        end

        for k,v in pairs(switch_but) do
            v:setSelected(false)
            v:setZOrder(v:getTag())
        end
        sender:setZOrder(100)
        sender:setSelected(true)
        getEffectOfGameSound(60)
    elseif eventType == 3  then
        sender:setSelected(not sender:isSelected())
    end
end

-- start --
--------------------------------
-- 以animation_config中配置的id获取动画数据
-- @function getAnimationByConfigId
-- @param number animationId 配置中的id
-- @param function asyncHandler 若希望异步加载,此参数为回调函数,参数为Animation动画数据,如已经加载,则下一帧直接回调
-- end --
function getAnimationByConfigId(animationId, asyncHandler)
    local scheduler = cc.Director:getInstance():getScheduler()
    local cfg = animation_config.animation_config[animationId]    
    local async = type(asyncHandler) == "function"
    local animation = display.getAnimationCache(cfg.name)
    if animation then
        if async then
            scheduler.performWithDelayGlobal(function()
                    asyncHandler(animation)
                end, 0)
        else
            return animation
        end
    else
    	local len = cfg.endFrame - cfg.startFrame + 1
        local time = cfg.duration / len
        if async then
            display.loadSpriteFrames(cfg.plist, cfg.png, function()
                local frames = display.newFrames(cfg.name .. "%02d.png", cfg.startFrame, len, cfg.direction ~= 0)
                animation = display.newAnimation(frames, time)
                display.setAnimationCache(cfg.name, animation)

                scheduler.performWithDelayGlobal(function()
                        asyncHandler(animation)
                    end, 0)
            end)
        else
            display.loadSpriteFrames(cfg.plist, cfg.png)
            local frames = display.newFrames(cfg.name .. "%02d.png", cfg.startFrame, len, cfg.direction ~= 0)
            animation = display.newAnimation(frames, time)
            display.setAnimationCache(cfg.name, animation)
            return animation
        end
    end
end

function getAnimationByConfigIdWithParams(animationId, fromFrame, toFrame, isReversed, time)
	local cfg = animation_config.animation_config[animationId]

	local animation = display.getAnimationCache(cfg.name)
	if not animation then
		display.loadSpriteFrames(cfg.plist, cfg.png)
	end
	local len = toFrame - fromFrame + 1
	local frames = display.newFrames(cfg.name .. "%02d.png", fromFrame, len, isReversed)
    animation = display.newAnimation(frames, time / len)
    display.setAnimationCache(cfg.name, animation)

    return animation
end


-- 中英混合截取
function GetShortName(sName,nMaxCount,nShowCount)
    if getStringSize(sName) > nMaxCount then 
        if sName == nil or nMaxCount == nil then
            return
        end
        local sStr = sName
        local tCode = {}
        local tName = {}
        local nLenInByte = #sStr
        local nWidth = 0
        if nShowCount == nil then
           nShowCount = nMaxCount - 3
        end
        for i=1,nLenInByte do
            local curByte = string.byte(sStr, i)
            local byteCount = 0;
            if curByte>0 and curByte<=127 then
                byteCount = 1
            elseif curByte>=192 and curByte<223 then
                byteCount = 2
            elseif curByte>=224 and curByte<239 then
                byteCount = 3
            elseif curByte>=240 and curByte<=247 then
                byteCount = 4
            end
            local char = nil
            if byteCount > 0 then
                char = string.sub(sStr, i, i+byteCount-1)
                i = i + byteCount -1
            end
            if byteCount == 1 then
                nWidth = nWidth + 1
                table.insert(tName,char)
                table.insert(tCode,1)
                
            elseif byteCount > 1 then
                nWidth = nWidth + 2
                table.insert(tName,char)
                table.insert(tCode,2)
            end
        end
        
        if nWidth > nMaxCount then
            local _sN = ""
            local _len = 0
            for i=1,#tName do
                _sN = _sN .. tName[i]
                _len = _len + tCode[i]
                if _len >= nShowCount then
                    break
                end
            end
            sName = _sN .. "..."
        end
    end
    return sName
end
-- 获取字符长度
function getStringSize(str)
    local num = #str
    local totalNum = 0
    for i=1,num do
        local curByte = string.byte(str, i)
        local byteCount = 0;
        if curByte>0 and curByte<=127 then
            byteCount = 1
        elseif curByte>=192 and curByte<223 then
            byteCount = 2
        elseif curByte>=224 and curByte<239 then
            byteCount = 2
        elseif curByte>=240 and curByte<=247 then
            byteCount = 2
        end
        totalNum = totalNum+byteCount
    end
    return totalNum
end

function dump_value_2(v)
    if type(v) == "string" then
        v = "\"" .. v .. "\""
    end
    return tostring(v)
end

function dumpToFile(value, desciption, nesting)
    if device.platform ~= "mac" or DEBUG == 0 then
        return
    end
    if type(nesting) ~= "number" then nesting = 3 end

    local lookupTable = {}
    local result = {}

    local traceback = string.split(debug.traceback("", 2), "\n")
    -- print("dump from: " .. string.trim(traceback[3]))

    local function dump_(value, desciption, indent, nest, keylen)
        desciption = desciption or "<var>"
        local spc = ""
        if type(keylen) == "number" then
            spc = string.rep(" ", keylen - string.len(dump_value_2(desciption)))
        end
        if type(value) ~= "table" then
            result[#result +1 ] = string.format("%s%s%s = %s", indent, dump_value_2(desciption), spc, dump_value_2(value))
        elseif lookupTable[tostring(value)] then
            result[#result +1 ] = string.format("%s%s%s = *REF*", indent, dump_value_2(desciption), spc)
        else
            lookupTable[tostring(value)] = true
            if nest > nesting then
                result[#result +1 ] = string.format("%s%s = *MAX NESTING*", indent, dump_value_2(desciption))
            else
                result[#result +1 ] = string.format("%s%s = {", indent, dump_value_2(desciption))
                local indent2 = indent.."    "
                local keys = {}
                local keylen = 0
                local values = {}
                for k, v in pairs(value) do
                    keys[#keys + 1] = k
                    local vk = dump_value_2(k)
                    local vkl = string.len(vk)
                    if vkl > keylen then keylen = vkl end
                    values[k] = v
                end
                table.sort(keys, function(a, b)
                    if type(a) == "number" and type(b) == "number" then
                        return a < b
                    else
                        return tostring(a) < tostring(b)
                    end
                end)
                for i, k in ipairs(keys) do
                    dump_(values[k], k, indent2, nest + 1, keylen)
                end
                result[#result +1] = string.format("%s}", indent)
            end
        end
    end
    dump_(value, desciption, "", 1)
    local stringALL = ""
    for i, line in ipairs(result) do
        stringALL = stringALL..line.."\n"
    end
    local outFileName = "" 
    if "报告表" == desciption then 
        local open_file = io.open(cc.FileUtils:getInstance():getWritablePath().."report.lua","a+")
        open_file:write(stringALL)
        open_file:flush()
        open_file:close()
    elseif "回放数据" == desciption then 
        local open_file = io.open(cc.FileUtils:getInstance():getWritablePath().."roomRpData.lua","w")
        open_file:write(stringALL)
        open_file:flush()
        open_file:close()
    else
        local open_file = io.open(cc.FileUtils:getInstance():getWritablePath().."roomData.lua","w")
        open_file:write(stringALL)
        open_file:flush()
        open_file:close()
    end
end

