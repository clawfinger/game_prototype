package game

import (
	gamecontext "prototype/game_context"
	"prototype/input"
	"prototype/state"
)

type GameInstance struct {
	states          *state.StateManager
	renderSystem    *RenderSystem
	InputController *input.InputController
}

func NewGameInstance() *GameInstance {
	return &GameInstance{
		states:          state.NewStateManager(),
		renderSystem:    NewRenderSystem(),
		InputController: input.NewInputController(gamecontext.GameContext.EventDispatcher),
	}
}

func (i *GameInstance) Init() {
	i.renderSystem.Init()
	i.states.PushState(state.NewGameLevelState()) // TODO: change to some kind of menu state
}

func (i *GameInstance) Update() {
	i.InputController.Update()
	i.states.GetCurrentState().Update()
}

func (i *GameInstance) Render() {
	i.states.GetCurrentState().Render()
	i.renderSystem.Render()
}
