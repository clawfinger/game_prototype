package systems

import (
	"prototype/ecs/components"
	"prototype/event"
	gamecontext "prototype/game_context"
	"prototype/screen"
)

type RenderSystem struct {
	System
	ed  *event.EventDispatcher
	ec  *components.EntityContainer
	scr *screen.Screen
}

func (s *RenderSystem) Update() {
}

func (s *RenderSystem) Render() {
	for _, entityID := range s.Entities {
		spriteComponent := components.GetComponent[*components.SpriteComponent](s.ec, entityID, components.SpriteComponentName)
		positionComponent := components.GetComponent[*components.PositionComponent](s.ec, entityID, components.PositionComponentName)
		s.scr.AddToLayer(spriteComponent.Layer, spriteComponent.Sprite, &positionComponent.Transform)
	}
}

func (s *RenderSystem) Init() {
	s.ed.Subscribe(event.EntityCreatedEventName, s)
	s.ed.Subscribe(event.RenderEventName, s)
}

func (s *RenderSystem) Notify(e event.Event) {
	switch event := e.(type) {
	case *event.EntityCreatedEvent:
		s.handleEntityCreatedEvent(event)
	case *event.RenderEvent:
		s.Render()
	}
}

func (s *RenderSystem) handleEntityCreatedEvent(e *event.EntityCreatedEvent) {
	if !s.HasEntity(e.EntityID) && s.FitsRequirements(e.Components) {
		s.System.Entities = append(s.System.Entities, e.EntityID)
	}
}

func NewRenderSystem() *RenderSystem {
	return &RenderSystem{
		System: System{
			Entities:     []int64{},
			Requirements: []string{components.PositionComponentName, components.SpriteComponentName},
		},
		ed:  gamecontext.GameContext.EventDispatcher,
		ec:  gamecontext.GameContext.EntityContainer,
		scr: gamecontext.GameContext.Screen,
	}
}
