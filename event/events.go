package event

import "reflect"

var (
	MousePressedEventName = reflect.TypeOf(MousePressedEvent{}).Name()
	MouseRelesedEventName = reflect.TypeOf(MouseRelesedEvent{}).Name()
	MouseMovedEventName   = reflect.TypeOf(MouseMovedEvent{}).Name()

	EntityCreatedEventName  = reflect.TypeOf(EntityCreatedEvent{}).Name()
	MapArenaLoadedEventName = reflect.TypeOf(MapArenaLoadedEvent{}).Name()
	RenderEventName         = reflect.TypeOf(RenderEvent{}).Name()
)

type MousePressedEvent struct {
	X int
	Y int
}

func (e *MousePressedEvent) Name() string {
	return MousePressedEventName
}

type MouseRelesedEvent struct {
	X int
	Y int
}

func (e *MouseRelesedEvent) Name() string {
	return MouseRelesedEventName
}

type MouseMovedEvent struct {
	X int
	Y int
}

func (e *MouseMovedEvent) Name() string {
	return MouseMovedEventName
}

type EntityCreatedEvent struct {
	EntityID   int64
	Components []string
}

func (e *EntityCreatedEvent) Name() string {
	return EntityCreatedEventName
}

type MapArenaLoadedEvent struct {
}

func (e *MapArenaLoadedEvent) Name() string {
	return MapArenaLoadedEventName
}

type RenderEvent struct {
}

func (e *RenderEvent) Name() string {
	return RenderEventName
}
