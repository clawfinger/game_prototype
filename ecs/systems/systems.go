package systems

import (
	"fmt"
	"prototype/event"
	gamecontext "prototype/game_context"
)

type SystemBase interface {
	Update()
	Init()
}

type System struct {
	Entities     []int64
	Requirements []string
}

func (s *System) FitsRequirements(entityComponents []string) bool {
	for _, requirement := range s.Requirements {
		contains := false
		for _, component := range entityComponents {
			if requirement == component {
				contains = true
				break
			}
		}
		if !contains {
			return false
		}
	}
	return true
}

func (s *System) HasEntity(entityID int64) bool {
	for _, entity := range s.Entities {
		if entity == entityID {
			return true
		}
	}
	return false
}

func (s *System) AddEntity(entityID int64) {
	if !s.HasEntity(entityID) {
		s.Entities = append(s.Entities, entityID)
	}
}

func (s *System) RemoveEntity(entityID int64) {
	for i, entity := range s.Entities {
		if entity == entityID {
			s.Entities[i], s.Entities[len(s.Entities)-1] = s.Entities[len(s.Entities)-1], s.Entities[i]
			s.Entities = s.Entities[:len(s.Entities)-1]
		}
	}
}

type MovementSystem struct {
	System
	ed *event.EventDispatcher
}

func (s *MovementSystem) Update() {

}

func NewMovementSystem() *MovementSystem {
	return &MovementSystem{
		ed: gamecontext.GameContext.EventDispatcher,
	}
}

func (s *MovementSystem) Init() {
	s.ed.Subscribe(event.EntityCreatedEventName, s)
	s.ed.Subscribe(event.MousePressedEventName, s)
	s.ed.Subscribe(event.MouseRelesedEventName, s)
	s.ed.Subscribe(event.MouseMovedEventName, s)
}

func (s *MovementSystem) Notify(e event.Event) {
	switch event := e.(type) {
	case *event.EntityCreatedEvent:
		s.handleEntityCreatedEvent(event)
	case *event.MousePressedEvent:
		s.handleMousePressedEvent(event)
	case *event.MouseRelesedEvent:
		s.handleMouseRelesedEvent(event)
	case *event.MouseMovedEvent:
		s.handleMouseMovedEvent(event)
	}
}

func (s *MovementSystem) handleEntityCreatedEvent(e *event.EntityCreatedEvent) {
	if !s.HasEntity(e.EntityID) && s.FitsRequirements(e.Components) {
		s.System.Entities = append(s.System.Entities, e.EntityID)
	}
}

func (s *MovementSystem) handleMouseRelesedEvent(e *event.MouseRelesedEvent) {
	fmt.Printf("mouse released at %d:%d\n", e.X, e.Y)
}

func (s *MovementSystem) handleMousePressedEvent(e *event.MousePressedEvent) {
	fmt.Printf("mouse pressed at %d:%d\n", e.X, e.Y)
}

func (s *MovementSystem) handleMouseMovedEvent(e *event.MouseMovedEvent) {
	fmt.Printf("mouse moved to %d:%d\n", e.X, e.Y)
}
