package tools

import (
	"../../common"
	"sync"
)

type ComponentCollecter struct {
	map4Component map[int]common.IComponent
}

var collecter *ComponentCollecter
var collecterOnce sync.Once

func GetComponentCollecter() *ComponentCollecter {
	collecterOnce.Do(func() {
		collecter = &ComponentCollecter{}
		collecter.init()
	})
	return collecter
}

func (o *ComponentCollecter) init() {
	o.map4Component = make(map[int]common.IComponent)
}

func (o *ComponentCollecter) AddComponent(identifier int, component common.IComponent) {
	o.map4Component[identifier] = component
}

func (o *ComponentCollecter) RemoveComponent(identifier int) (common.IComponent, bool) {
	v, ok := o.map4Component[identifier]
	delete(o.map4Component, identifier)

	return v, ok
}

func (o *ComponentCollecter) StopComponent(identifier int) {
	v, ok := o.RemoveComponent(identifier)
	if ok {
		v.Stop()
	}
}

func (o *ComponentCollecter) GetComponent(identifier int) (common.IComponent, bool) {
	v, ok := o.map4Component[identifier]
	return v, ok
}