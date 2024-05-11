package game

import "prototype/state"

type GameInstance struct {
	states       *state.StateManager
	renderSystem *RenderSystem
}

func NewGameInstance() *GameInstance {
	return &GameInstance{
		states:       state.NewStateManager(),
		renderSystem: NewRenderSystem(),
	}
}

func (i *GameInstance) Init() {
	i.states.PushState(state.NewGameLevelState()) // TODO: change to some kind of menu state
	i.renderSystem.Init()
}

func (i *GameInstance) Update() {
	i.states.GetCurrentState().Update()
}

func (i *GameInstance) Render() {
	i.states.GetCurrentState().Render()
	i.renderSystem.Render()
}
