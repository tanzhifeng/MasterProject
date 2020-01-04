package logic

import (
	"../../common"
	"../include"
	"fmt"
	"sync"
)

type LogicalEvent struct {
	Event int
	Options *[]interface{}
}

type Logical struct {
	channelStop chan bool
	channel4Loop chan *LogicalEvent
	map4LogicalHandler map[int]include.LogicalHandler
}

var once sync.Once
var instance *Logical

func GetLogical() *Logical {
	once.Do(func() {
		instance = &Logical{}
		instance.init()
	})
	return instance
}

func (o *Logical) init() {
	o.map4LogicalHandler = make(map[int]include.LogicalHandler)

	o.map4LogicalHandler[common.EVENT_LOGICAL_CLIENT_CONNECTED] = o.LogicalClientConnected
	o.map4LogicalHandler[common.EVENT_LOGICAL_CLIENT_LOSTED] = o.LogicalClientLosted

	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_CONNECTED] = o.LoginServerClientConnected
	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_LOSTED] = o.LoginServerClientLost
	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_CHECK_ALIVE] = o.LoginServerClientCheckAlive
	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_BEATS] = o.LoginServerClientBeats
	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_LOGIN] = o.LoginServerUserLogin
	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_LOGIN_RESULT] = o.LoginServerUserLoginResult
	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_REGISTER] = o.LoginServerUserRegister
	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_REGISTER_RESULT] = o.LoginServerUserRegisterResult
	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_USER_ENTER_ROOM] = o.LoginServerUserEnterRoom
	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_USER_ENTER_ROOM_RESULT] = o.LoginServerUserEnterRoomResult
	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_USER_LEAVE_ROOM] = o.LoginServerUserLeaveRoom
	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_USER_LEAVE_ROOM_RESULT] = o.LoginServerUserLeaveRoomResult

	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_NOTIFY_APPEND_ROOM] = o.LoginServerNotifyAppendRoom
	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_NOTIFY_REMOVE_ROOM] = o.LoginServerNotifyRemoveRoom

	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_NOTIFY_APPEND_ROOM_USER] = o.LoginServerNotifyAppendRoomUser
	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_NOTIFY_REMOVE_ROOM_USER] = o.LoginServerNotifyRemoveRoomUser
}

func (o *Logical) Start() bool {
	o.channelStop = make(chan bool)
	o.channel4Loop = make(chan *LogicalEvent, common.DefaultChannelLen)

	go o.Loop()

	return true
}

func (o *Logical) Stop() {
	close(o.channel4Loop)

	<-o.channelStop
}

func (o *Logical) AppendEvent(event int, options ...interface{}) {
	o.channel4Loop <- &LogicalEvent{Event:event, Options:&options}
}

func (o *Logical) HandleEvent(event *LogicalEvent) {
	v, ok := o.map4LogicalHandler[event.Event]
	if ok {
		v(*event.Options...)
	} else {
		fmt.Printf("Unknow Logical Event [%d]!\n", event.Event)
	}
}

func (o *Logical) Loop() {
	defer close(o.channelStop)

	for {
		v, ok := <-o.channel4Loop
		if !ok {
			return
		}

		o.HandleEvent(v)
	}
}