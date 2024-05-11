package state

import gamemap "prototype/map"

type GameLevelState struct {
	Map *gamemap.GameMap
}

func NewGameLevelState() *GameLevelState {
	return &GameLevelState{
		Map: gamemap.NewMap(),
	}
}

func (s *GameLevelState) Init() {
	s.Map.Init()
}

func (s *GameLevelState) Update() {
}

func (s *GameLevelState) Render() {
	s.Map.Render()
}

func (s *GameLevelState) Deinit() {}
