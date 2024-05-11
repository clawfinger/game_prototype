package state

type IState interface {
	Init()
	Update()
	Render()
	Deinit()
}

type StateManager struct {
	states []IState
}

func NewStateManager() *StateManager {
	return &StateManager{
		states: make([]IState, 0),
	}
}

func (s *StateManager) Update() {
	for i := range s.states {
		s.states[i].Update()
	}
}

func (s *StateManager) PushState(state IState) {
	state.Init()
	s.states = append(s.states, state)
}

func (s *StateManager) PopState() {
	state := s.states[len(s.states)-1]
	state.Deinit()
	s.states = s.states[:len(s.states)-1]
}

func (s *StateManager) GetCurrentState() IState {
	return s.states[len(s.states)-1]
}
