package loginsvr

import (
	"time"

	"../../common"
	"../../common/tools"
	"../logic"
)

func (o *LoginServer) checkBeatsStart() {
	o.scheduleBeats = tools.GetScheduler().Schedule(common.BeatsIntervalServer * time.Second, tools.LoopForever, o.checkClientAlive)
}

func (o *LoginServer) checkClientAlive(params ...interface{}) {
	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_CHECK_ALIVE)
}

func (o *LoginServer) checkBeatsStop() {
	tools.GetScheduler().Unschedule(o.scheduleBeats)
}

func (o *LoginServer) onLoginClientLost(socket uint32, bytes []byte) {
	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_LOSTED, socket)
}

func (o *LoginServer) onLoginClientBeats(socket uint32, bytes []byte) {
	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_BEATS, socket)
}