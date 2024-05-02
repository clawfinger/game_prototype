package ecs

import (
	"prototype/event"
	"prototype/screen"
)

type SystemBase interface {
	Update()
	Init()
}

type System struct {
	entities     []int64
	requirements []string
}

func (s *System) fitsRequirements(entityComponents []string) bool {
	for _, requirement := range s.requirements {
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

func (s *System) hasEntity(entityID int64) bool {
	for _, entity := range s.entities {
		if entity == entityID {
			return true
		}
	}
	return false
}

func (s *System) addEntity(entityID int64) {
	if !s.hasEntity(entityID) {
		s.entities = append(s.entities, entityID)
	}
}

func (s *System) removeEntity(entityID int64) {
	for i, entity := range s.entities {
		if entity == entityID {
			s.entities[i], s.entities[len(s.entities)-1] = s.entities[len(s.entities)-1], s.entities[i]
			s.entities = s.entities[:len(s.entities)-1]
		}
	}
}

type MovementSystem struct {
	System
}

func (s *MovementSystem) Update() {

}

func (s *MovementSystem) Init() {

}

func NewMovementSystem() *MovementSystem {
	return &MovementSystem{}
}

type RenderSystem struct {
	System
	ed  *event.EventDispatcher
	ec  *EntityContainer
	scr *screen.Screen
}

func (s *RenderSystem) Update() {
	for _, entityID := range s.entities {
		spriteComponent, err := GetComponent[*SpriteComponent](s.ec, entityID, SpriteComponentName)
		if err != nil {
			// TODO: log
			continue
		}
		positionComponent, err := GetComponent[*PositionComponent](s.ec, entityID, PositionComponentName)
		if err != nil {
			// TODO: log
			continue
		}
		s.scr.AddToLayer(screen.ActorsLayer, spriteComponent.Sprite, &positionComponent.transform)
	}
}

func (s *RenderSystem) Init() {
	s.ed.Subscribe(event.EntityCreatedEventName, s)
}

func (s *RenderSystem) Notify(e event.Event) {
	switch event := e.(type) {
	case *event.EntityCreatedEvent:
		s.handleEntityCreatedEvent(event)
	}
}

func (s *RenderSystem) handleEntityCreatedEvent(e *event.EntityCreatedEvent) {
	if !s.hasEntity(e.EntityID) {
		s.System.entities = append(s.System.entities, e.EntityID)
	}
}

func NewRenderSystem(ed *event.EventDispatcher, ec *EntityContainer, s *screen.Screen) *RenderSystem {
	return &RenderSystem{
		System: System{
			entities:     []int64{},
			requirements: []string{PositionComponentName, SpriteComponentName},
		},
		ed:  ed,
		ec:  ec,
		scr: s,
	}
}
