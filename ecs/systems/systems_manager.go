package systems

import (
	"prototype/ecs/components"
	"prototype/event"
	"prototype/screen"
)

type SystemManager struct {
	systems []SystemBase
}

func NewSystemManager(ed *event.EventDispatcher, ec *components.EntityContainer, s *screen.Screen) *SystemManager {
	sm := &SystemManager{}
	sm.systems = append(sm.systems,
		NewMovementSystem(),
		NewRenderSystem())
	return sm
}

func (s *SystemManager) Init() {
	for _, system := range s.systems {
		system.Init()
	}
}

func (s *SystemManager) Update() {
	for _, system := range s.systems {
		system.Update()
	}
}
