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
	config *map[string]interface{}
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

	o.map4LogicalHandler[common.EVENT_ROOM_SERVER_CLIENT_CONNECTED] = o.RoomServerClientConnected
	o.map4LogicalHandler[common.EVENT_ROOM_SERVER_CLIENT_LOSTED] = o.RoomServerClientLost
}

func (o *Logical) GetConfig() *map[string]interface{} {
	return o.config
}

func (o *Logical) Start(config *map[string]interface{}) bool {
	o.config = config
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