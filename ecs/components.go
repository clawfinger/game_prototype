package ecs

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	PositionComponentName = reflect.TypeOf(PositionComponent{}).Name()
	SpriteComponentName   = reflect.TypeOf(SpriteComponent{}).Name()
)

type ComponentBase interface {
	Name() string
}

type PositionComponent struct {
	X         int
	Y         int
	transform ebiten.GeoM
}

func (c *PositionComponent) Name() string {
	return PositionComponentName
}

type SpriteComponent struct {
	Sprite *ebiten.Image
}

func (c *SpriteComponent) Name() string {
	return SpriteComponentName
}
