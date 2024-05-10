package ecs

import (
	"prototype/ecs/components"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestECS(t *testing.T) {
	ec := components.NewEntityContainer()
	id, err := ec.CreateEntity(components.StringList{components.PositionComponentName, components.SpriteComponentName})
	assert.NoError(t, err)
	component, err := components.GetComponent[*components.PositionComponent](ec, id, components.PositionComponentName)
	assert.NoError(t, err)
	component.X = 666
	component.Y = 999
	componentN, err := components.GetComponent[*components.PositionComponent](ec, id, components.PositionComponentName)
	assert.NoError(t, err)
	assert.Equal(t, 666, componentN.X)
	assert.Equal(t, 999, componentN.Y)
}
