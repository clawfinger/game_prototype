package ecs

import "reflect"

var (
	PositionComponentName = reflect.TypeOf(PositionComponent{}).Name()
)

type ComponentBase interface {
	Name() string
}

type PositionComponent struct {
	X int
	Y int
}

func (c *PositionComponent) Name() string {
	return PositionComponentName
}
