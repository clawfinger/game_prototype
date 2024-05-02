package ecs

import "fmt"

type StringList []string

type EntityContainer struct {
	currentEntityID    int64
	components         map[int64]map[string]ComponentBase
	componentFactories map[string]func() ComponentBase
}

func NewEntityContainer() *EntityContainer {
	ec := &EntityContainer{
		components:         map[int64]map[string]ComponentBase{},
		componentFactories: map[string]func() ComponentBase{},
	}
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
	c.componentFactories[SpriteComponentName] = func() ComponentBase {
		return &SpriteComponent{}
	}
}

func GetComponent[T any](c *EntityContainer, entityID int64, name string) (T, error) {
	components, ok := c.components[entityID]
	var t T
	if !ok {
		return t, fmt.Errorf("no entity %d", entityID)
	}
	component, ok := components[name]
	if !ok {
		return t, fmt.Errorf("no component %s for entity %d", name, entityID)
	}
	return component.(T), nil
}
