package components

import (
	"fmt"
	"prototype/event"
)

type EntityContainer struct {
	currentEntityID    int64
	components         map[int64]map[string]ComponentBase
	componentFactories map[string]func() ComponentBase
	ed                 *event.EventDispatcher
}

func NewEntityContainer(ed *event.EventDispatcher) *EntityContainer {
	ec := &EntityContainer{
		components:         map[int64]map[string]ComponentBase{},
		componentFactories: map[string]func() ComponentBase{},
		ed:                 ed,
	}
	ec.createComponentFactories()
	return ec
}

func (c *EntityContainer) CreateEntity(components []string) int64 {
	entityData := make(map[string]ComponentBase)
	for _, component := range components {
		factory, ok := c.componentFactories[component]
		if !ok {
			// TODO: log
		}
		entityData[component] = factory()
	}
	entityID := c.currentEntityID
	c.components[entityID] = entityData
	c.currentEntityID++
	c.ed.Dispatch(&event.EntityCreatedEvent{
		EntityID:   entityID,
		Components: components,
	})
	return entityID
}

func (c *EntityContainer) RemoveEntity(entityID int64) {
	delete(c.components, entityID)
}

func (c *EntityContainer) HasComponent(entityID int64, component string) (bool, error) {
	entityData, ok := c.components[entityID]
	if !ok {
		return false, fmt.Errorf("no data for entity %d", entityID)
	}
	_, ok = entityData[component]
	return ok, nil
}

func (c *EntityContainer) createComponentFactories() {
	c.componentFactories[PositionComponentName] = func() ComponentBase {
		return &PositionComponent{}
	}
	c.componentFactories[SpriteComponentName] = func() ComponentBase {
		return &SpriteComponent{}
	}
}

func GetComponent[T any](c *EntityContainer, entityID int64, name string) T {
	components, ok := c.components[entityID]
	if !ok {
		// TODO log
	}
	component, ok := components[name]
	if !ok {
		// TODO log
	}
	return component.(T)
}
