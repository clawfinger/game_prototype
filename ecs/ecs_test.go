package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestECS(t *testing.T) {
	ec := NewEntityContainer()
	id, err := ec.CreateEntity(StringList{PositionComponentName, SpriteComponentName})
	assert.NoError(t, err)
	component, err := GetComponent[*PositionComponent](ec, id, PositionComponentName)
	assert.NoError(t, err)
	component.X = 666
	component.Y = 999
	componentN, err := GetComponent[*PositionComponent](ec, id, PositionComponentName)
	assert.NoError(t, err)
	assert.Equal(t, 666, componentN.X)
	assert.Equal(t, 999, componentN.Y)
}
