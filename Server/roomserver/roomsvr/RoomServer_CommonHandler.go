package roomsvr

import (
	"../../common"
	"../logic"
)

func (o *RoomServer) onLoginClientLost(socket uint32, bytes []byte) {
	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_LOSTED, socket)
}

func (o *RoomServer) onLoginClientBeats(socket uint32, bytes []byte) {
	logic.GetLogical().AppendEvent(common.EVENT_LOGIN_SERVER_CLIENT_BEATS, socket)
}