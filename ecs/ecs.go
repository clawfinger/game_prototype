package ecs

import (
	"fmt"
)

type StringList []string

type EntityContainer struct {
	currentEntityID    int64
	components         map[int64]map[string]ComponentBase
	componentFactories map[string]func() ComponentBase
}

func NewEntityContainer() *EntityContainer {
	ec := &EntityContainer{}
	ec.createComponentFactories()
	return ec
}

func (c *EntityContainer) CreateEntity(components StringList) (int64, error) {
	entityData := make(map[string]ComponentBase)
	for _, component := range components {
		factory, ok := c.componentFactories[component]
		if !ok {
			return 0, fmt.Errorf("no factory for component %s", component)
		}
		entityData[component] = factory()
	}
	entityID := c.currentEntityID
	c.components[entityID] = entityData
	c.currentEntityID++
	return entityID, nil
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
}
