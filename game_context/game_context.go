package gamecontext

import (
	"prototype/ecs/components"
	"prototype/event"
	"prototype/screen"
	"prototype/tiles"
)

var GameContext *GameContextData

type GameContextData struct {
	EventDispatcher *event.EventDispatcher
	EntityContainer *components.EntityContainer
	Screen          *screen.Screen
	TileManager     *tiles.TileManager
}

func NewGameContext() *GameContextData {
	ed := event.NewEventDispatcher()
	return &GameContextData{
		EventDispatcher: ed,
		EntityContainer: components.NewEntityContainer(ed),
		Screen:          screen.NewScreen(),
		TileManager:     tiles.NewTileManager(),
	}
}

func Init() {
	GameContext = NewGameContext()
	GameContext.TileManager.Init()
}
