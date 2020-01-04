package tools

import (
	"sync"
	"time"
)

//LoopForever ...
const LoopForever = ^uint64(0)

var scheduler *Scheduler
var schedulerOnce sync.Once
var scheduleSeed uint

//HandlerSchedule ...
type HandlerSchedule func(params ...interface{})

//ScheduleData ...
type ScheduleData struct {
	TID     uint
	T       *time.Ticker
	Repeat  uint64
	Params  []interface{}
	Handler HandlerSchedule
}

//Scheduler ..
type Scheduler struct {
	MapSchedule     map[uint]*ScheduleData
}

//UniqueID ...
func UniqueID() uint {
	scheduleSeed++

	return scheduleSeed
}

//GetScheduler ...
func GetScheduler() *Scheduler {
	schedulerOnce.Do(func() {
		scheduler = &Scheduler{}
		scheduler.init()
	})
	return scheduler
}

func (o *Scheduler) init() {
	o.MapSchedule = make(map[uint]*ScheduleData)
}

//Stop ...
func (o *Scheduler) Stop() {
	for _, v := range o.MapSchedule {
		v.T.Stop()
	}
	o.MapSchedule = nil
}

//Schedule ...
func (o *Scheduler) Schedule(d time.Duration, repeat uint64, handler HandlerSchedule, params ...interface{}) uint {
	schedule := &ScheduleData{
		TID:     UniqueID(),
		T:       time.NewTicker(d),
		Repeat:  repeat,
		Params:  params,
		Handler: handler,
	}

	o.MapSchedule[schedule.TID] = schedule

	go onSchedule(o, schedule.TID)

	return schedule.TID
}

func onSchedule(o *Scheduler, unique uint) {
	schedule := o.MapSchedule[unique]
	if schedule != nil {
		for {
			<-schedule.T.C

			schedule.Repeat--

			schedule.Handler(schedule.Params...)

			if schedule.Repeat <= 0 {
				GetScheduler().Unschedule(schedule.TID)
				return
			}
		}
	}
}

//Unschedule ...
func (o *Scheduler) Unschedule(unique uint) {
	schedule, ok := o.MapSchedule[unique]
	if ok {
		schedule.T.Stop()
		delete(o.MapSchedule, unique)
	}
}