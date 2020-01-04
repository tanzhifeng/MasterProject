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

	o.map4LogicalHandler[common.EVENT_DATABASE_SERVER_CLIENT_CONNECTED] = o.LogicalServerConnected

	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_LOSTED] = o.LogicalUserLosted
	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_LOGIN] = o.LogicalUserLogin
	o.map4LogicalHandler[common.EVENT_LOGIN_SERVER_CLIENT_REGISTER] = o.LogicalUserRegister
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