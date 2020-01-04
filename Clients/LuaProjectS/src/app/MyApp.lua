--配置文件
require "configs.cfg"
-- 全局文件和工具类
require "global.global"
require "util.tools"
require "util.common"
require "util.serializable"
--协议文件
require "handler.gt_login_handler"
require "handler.gt_room_handler"
require "handler.gatesvr"
--数据文件
require "Data.PlayerData"

--UI文件
require "app.views.init"

local MyApp = class("MyApp", cc.load("mvc").AppBase)

function MyApp:onCreate()
    math.randomseed(os.time())
    global.g_gameDispatcher = require("util.GameDispatcher").new()
    global.g_playerData = PlayerData.new()
end

return MyApp
