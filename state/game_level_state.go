package state

import (
	"prototype/ecs/components"
	"prototype/ecs/systems"
	gamecontext "prototype/game_context"
	gamemap "prototype/map"
	"prototype/screen"
)

type GameLevelState struct {
	Map               *gamemap.GameMap
	MapRenderer       *gamemap.MapRenderer
	EntityContainer   *components.EntityContainer
	systemsManager    *systems.SystemManager
	characterEntityID int64
}

func NewGameLevelState() *GameLevelState {
	gameMap := gamemap.NewMap()
	return &GameLevelState{
		Map:             gameMap,
		MapRenderer:     gamemap.NewMapRenderer(gameMap),
		EntityContainer: gamecontext.GameContext.EntityContainer,
		systemsManager:  systems.NewSystemManager(gamecontext.GameContext.EventDispatcher, gamecontext.GameContext.EntityContainer, gamecontext.GameContext.Screen),
	}
}

func (s *GameLevelState) Init() {
	s.Map.Init()
	s.MapRenderer.Init()
	s.systemsManager.Init()
	s.characterEntityID = s.EntityContainer.CreateEntity([]string{components.PositionComponentName, components.SpriteComponentName})
	sprite := components.GetComponent[*components.SpriteComponent](s.EntityContainer, s.characterEntityID, components.SpriteComponentName)
	sprite.Layer = screen.ActorsLayer
	sprite.Sprite = gamecontext.GameContext.TileManager.GetTile(468)

	position := components.GetComponent[*components.PositionComponent](s.EntityContainer, s.characterEntityID, components.PositionComponentName)
	position.Transform.Translate(80+16, 40+16)
}

func (s *GameLevelState) Update() {
	s.systemsManager.Update()
}

func (s *GameLevelState) Render() {
	s.MapRenderer.Render()
}

func (s *GameLevelState) Deinit() {}
