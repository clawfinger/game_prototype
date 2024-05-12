package ecs

import (
	"prototype/ecs/components"
	"prototype/event"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestECS(t *testing.T) {
	ec := components.NewEntityContainer(event.NewEventDispatcher())
	id := ec.CreateEntity([]string{components.PositionComponentName, components.SpriteComponentName})
	component := components.GetComponent[*components.PositionComponent](ec, id, components.PositionComponentName)
	component.X = 666
	component.Y = 999
	componentN := components.GetComponent[*components.PositionComponent](ec, id, components.PositionComponentName)
	assert.Equal(t, 666, componentN.X)
	assert.Equal(t, 999, componentN.Y)
}
