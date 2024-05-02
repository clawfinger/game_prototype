package event

import "reflect"

var (
	MousePressedEventName  = reflect.TypeOf(MousePressedEvent{}).Name()
	EntityCreatedEventName = reflect.TypeOf(EntityCreatedEvent{}).Name()
)

type MousePressedEvent struct {
	x int
	y int
}

func (e *MousePressedEvent) Name() string {
	return MousePressedEventName
}

type EntityCreatedEvent struct {
	EntityID  int64
	Component []string
}

func (e *EntityCreatedEvent) Name() string {
	return MousePressedEventName
}
