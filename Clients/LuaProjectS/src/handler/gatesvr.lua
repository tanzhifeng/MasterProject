module("gatesvr", package.seeall)

local fileUtils = cc.FileUtils:getInstance()

g_gt_init = g_gt_init or false
g_gt_sock = g_gt_sock or clientSock.new()
g_gt_conn = g_gt_conn or false
g_gt_tick = g_gt_tick or nil

protoc_files = {
    "prot/pt_com/common.proto",
    "prot/pt_login/constants.proto",
    "prot/pt_login/login.proto",
    "prot/pt_login/room.proto",
}

for i = 1, #protoc_files do
    local fname = protoc_files[i]
    local path = fileUtils:fullPathForFilename(fname)
    local data = pb.io.read(path)
    local ret, pos = pb.load(data)
    if not ret then
        local err = string.format("load file %s failed at offset %s !", fname, pos)
        printf(err)
    end
end

pid = {}
for name, number, type in pb.fields("pt_login.login_pid") do
    print("名字是:",name,number)
    pid[name] = number
end
print("内容是：",pid.GT_LOGIN)
codes = {}
for name, number, type in pb.fields("pt_login.login_code") do
    codes[name] = number
end

otime = otime or nil

function adjustTime(time)
    if otime then
        os.time = otime
    end

    local nowTime = os.time()
    local delta = nowTime - time
    if math.abs(delta) > 5 then
        otime = os.time
        os.time = function(...)
                local o = otime(...)
                return o - delta
            end
    end
end

--
-- 网络相关
--
-- 连接gateserver

function startGate()
    if g_gt_init then return end

    g_gt_init = true
    
    if not g_gt_conn then
        g_gt_sock:connectByIp(GATE_SERV_IP, GATE_SERV_PORT, handlerReply)

        checkUpdateStart()
    end
end


function closeGate()
    if not g_gt_init then return end

    g_gt_init = false
    g_gt_conn = false
    beatsStop()
    checkUpdateStop()

    g_gt_sock:closeSocket()
end

function checkUpdateStart()
    checkUpdateStop()

    g_gt_tick = cc.Director:getInstance():getScheduler():scheduleScriptFunc(function()
            g_gt_sock:checkUpdate()
        end, 0, false)
end

function checkUpdateStop()
    if g_gt_tick then
        cc.Director:getInstance():getScheduler():unscheduleScriptEntry(g_gt_tick)
        g_gt_tick = nil
    end
end

g_gt_beats = g_gt_beats or nil

function beatsStart()
    beatsStop()

    g_gt_beats = cc.Director:getInstance():getScheduler():scheduleScriptFunc(function()
            sendGtData(pid.GT_BEATS)
        end, 10, false)
end

function beatsStop()
    if g_gt_beats then
        cc.Director:getInstance():getScheduler():unscheduleScriptEntry(g_gt_beats)
        g_gt_beats = nil
    end
end

function handlerReply(status, ptid, buffer)
    local fun = fun_tbl[status]
    fun(ptid, buffer)
end

-- gt
function sendGtData(ptid, data)
    if not g_gt_conn then return end

    local handlerData = handlerNetwork[ptid]

    if handlerData and handlerData.protocSend then
        local bytes = assert(pb.encode(handlerData.protocSend, data))
        g_gt_sock:handlerSend(ptid, bytes)
    else
        g_gt_sock:handlerSend(ptid)
    end

    print("in sendGtData Ptid:%d", ptid)

    return true
end

function onGateConnsucc()
    -- 连接成功
    g_gt_conn = true
    print("onGateConnsucc ip: %s port: %d", GATE_SERV_IP, GATE_SERV_PORT)
    global.g_gameDispatcher:dispatchEvent(EVENT_CONNECT_SUCCESSED)
    --连接成功后自动登录
    local account = global.g_playerData:getAccount()
    local password = global.g_playerData:getPassword()
    sendLogin(account,password)
end

function onGateConnfail()
    TextTipsUtils:showTips("连接服务器失败，请重试！")
    -- SwitchViewBase.switchView(SWITCH_VIEW_GROUP_ONE,1,HallViewPanel)
    -- 连接失败
    g_gt_init = false
    g_gt_conn = false
    print("onGateConnfail ip: %s port: %d", GATE_SERV_IP, GATE_SERV_PORT)

    checkUpdateStop()

    global.g_gameDispatcher:dispatchEvent(EVENT_CONNECT_FAILED)
end

function onGateConnlost()
    g_gt_init = false
    g_gt_conn = false
    print("onGateConnlost ip: %s port: %d", GATE_SERV_IP, GATE_SERV_PORT)

    beatsStop()
    checkUpdateStop()

    global.g_gameDispatcher:dispatchEvent(EVENT_CONNECT_LOSTED)
end

function onGateData(ptid, buffer)
    print(" onGateData ip: %s port: %d ptid:%d", GATE_SERV_IP, GATE_SERV_PORT, ptid)

    local handlerData = handlerNetwork[ptid]
    if not handlerData then
        print("in onGateData unmatch ptid:%d", ptid)
    else
        if handlerData.protocRecv then
            local data = assert(pb.decode(handlerData.protocRecv, buffer))
            handlerData.handler(data)
        else
            handlerData.handler()
        end
    end
end

fun_tbl = {
    [ clientSock.Succeed ] = onGateConnsucc,
    [ clientSock.Failed ] = onGateConnfail,
    [ clientSock.Losted ] = onGateConnlost,
    [ clientSock.Dataed ] = onGateData,
}

handlerNetwork = {
    [pid.LOGIN_LOGIN] = {
        handler = onLogin,
        protocSend = "pt_login.c2s_login_user",
        protocRecv = "pt_login.s2c_login_user_result",
    },
    [pid.LOGIN_REGISTER] = {
        handler = onRegister,
        protocSend = "pt_login.c2s_register_user",
        protocRecv = "pt_login.s2c_register_user_result",
    },
    -- [pid.GT_ROOM_NOTIFY_APPEND_ROOM] = {
    --     handler = onNotifyRoomAdded,
    --     protocRecv = "pt_gate.s2c_notify_append_room",
    -- },
    -- [pid.GT_ROOM_NOTIFY_REMOVE_ROOM] = {
    --     handler = onNotifyRoomRemoved,
    --     protocRecv = "pt_gate.s2c_notify_remove_room",
    -- },

    -- [pid.GT_ROOM_NOTIFY_APPEND_TABLE] = {
    --     handler = onNotifyTableAdded,
    --     protocRecv = "pt_gate.s2c_notify_append_table",
    -- },
    -- [pid.GT_ROOM_NOTIFY_REMOVE_TABLE] = {
    --     handler = onNotifyTableRemoved,
    --     protocRecv = "pt_gate.s2c_notify_remove_table",
    -- },

    -- [pid.GT_ROOM_USER_ENTER] = {
    --     handler = onRoomUserEnter,
    --     protocSend = "pt_gate.c2s_room_user_enter",
    --     protocRecv = "pt_gate.s2c_room_user_enter_result",
    -- },
    -- [pid.GT_ROOM_USER_LEAVE] = {
    --     handler = onRoomUserLeave,
    --     protocRecv = "pt_gate.s2c_room_user_leave_result",
    -- },
    -- [pid.GT_ROOM_NOTIFY_APPEND_USER] = {
    --     handler = onNotifyRoomAppendUser,
    --     protocRecv = "pt_gate.s2c_room_notify_append_user",
    -- },
    -- [pid.GT_ROOM_NOTIFY_REMOVE_USER] = {
    --     handler = onNotifyRoomRemoveUser,
    --     protocRecv = "pt_gate.s2c_room_notify_remove_user",
    -- },
    -- [pid.GT_ROOM_USER_SITDOWN] = {
    --     handler = onRoomUserSitdown,
    --     protocSend = "pt_gate.c2s_room_user_sitdown",
    --     protocRecv = "pt_gate.s2c_room_user_sitdown_result",
    -- },
    -- [pid.GT_ROOM_USER_STANDUP] = {
    --     handler = onRoomUserStandup,
    --     protocRecv = "pt_gate.s2c_room_user_standup_result",
    -- },
    -- [pid.GT_ROOM_NOTIFY_USER_SITDOWN] = {
    --     handler = onRoomNotifyUserSitdown,
    --     protocRecv = "pt_gate.s2c_room_notify_user_sitdown",
    -- },
    -- [pid.GT_ROOM_NOTIFY_USER_STANDUP] = {
    --     handler = onRoomNotifyUserStandup,
    --     protocRecv = "pt_gate.s2c_room_notify_user_standup",
    -- },
}