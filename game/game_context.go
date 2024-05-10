package game

import (
	"prototype/ecs/components"
	"prototype/event"
	"prototype/screen"
)

type GameContext struct {
	EventDispatcher *event.EventDispatcher
	EntityContainer *components.EntityContainer
	Screen          screen.Screen
}

func NewGameContext() *GameContext {
	return &GameContext{
		EventDispatcher: event.NewEventDispatcher(),
		EntityContainer: components.NewEntityContainer(),
		Screen:          *screen.NewScreen(),
	}
}
