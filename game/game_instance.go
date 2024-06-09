package game

import (
	"prototype/input"
	"prototype/state"
)

type GameInstance struct {
	states          *state.StateManager
	InputController *input.InputController
}

func NewGameInstance() *GameInstance {
	return &GameInstance{
		states:          state.NewStateManager(),
		InputController: input.NewInputController(),
	}
}

func (i *GameInstance) Init() {
	i.states.PushState(state.NewGameLevelState()) // TODO: change to some kind of menu state
}

func (i *GameInstance) Update() {
	i.InputController.Update()
	i.states.GetCurrentState().Update()
}

func (i *GameInstance) Render() {
	i.states.GetCurrentState().Render()
}
