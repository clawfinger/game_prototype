package gamemap

import (
	"fmt"
	gamecontext "prototype/game_context"
	"prototype/settings"
	"testing"
)

func TestLinear(t *testing.T) {
	m := GameMap{
		Settings: &MapSettings{},
	}
	m.Settings.SizeHorizontal = 5
	m.Settings.SizeVertical = 7
	idxMap := []int{
		7, 1, 1, 1, 1,
		1, 5, 2, 2, 1,
		1, 2, 2, 2, 1,
		1, 2, 2, 2, 1,
		1, 2, 2, 2, 1,
		1, 2, 2, 2, 1,
		1, 1, 1, 1, 9,
	}
	var tiles []*Tile
	for i := range idxMap {
		tiles = append(tiles, &Tile{SpriteIdx: idxMap[i]})
	}
	m.Settings.Map = tiles
	tests := []struct {
		name     string
		row      int
		column   int
		expected int
	}{
		{
			name:     "simple",
			row:      1,
			column:   1,
			expected: 5,
		},
		{
			name:     "first",
			row:      0,
			column:   0,
			expected: 7,
		},
		{
			name:     "last",
			row:      6,
			column:   4,
			expected: 9,
		},
	}
	for _, test := range tests {
		if res := m.createLinearFromRowAndColumn(test.row, test.column); m.Settings.Map[res].SpriteIdx != test.expected {
			t.Errorf(fmt.Sprintf("%s failed, expected %d, got %d, index %d", test.name, test.expected, m.Settings.Map[res], res))
		}
	}
}

func TestGetTiles(t *testing.T) {
	settings.Init()
	gamecontext.Init()
	m := NewMap()
	m.Init()
	m.Settings.SizeHorizontal = 5
	m.Settings.SizeVertical = 7
	idxMap := []int{
		7, 1, 1, 1, 1,
		1, 99, 2, 2, 1,
		1, 2, 2, 2, 1,
		1, 2, 2, 2, 1,
		1, 2, 2, 2, 1,
		1, 2, 2, 2, 1,
		1, 1, 1, 1, 9,
	}
	var tiles []*Tile
	for i := range idxMap {
		tiles = append(tiles, &Tile{SpriteIdx: idxMap[i]})
	}
	m.Settings.Map = tiles
	tests := []struct {
		name          string
		clickX        int
		clickY        int
		expectedTile  int
		expectedFound bool
	}{
		{
			name:          "not found",
			clickX:        10000,
			clickY:        10000,
			expectedFound: false,
		},
		{
			name:          "found",
			clickX:        16 + 8,
			clickY:        16 + 8,
			expectedTile:  99,
			expectedFound: true,
		},
		{
			name:          "found between",
			clickX:        16,
			clickY:        16,
			expectedTile:  99,
			expectedFound: true,
		},
	}
	for _, test := range tests {
		res, ok := m.GetTile(test.clickX, test.clickY)
		if !test.expectedFound {
			if ok != test.expectedFound {
				t.Errorf("not found not es axpected")
			}
			continue
		}

		if res.SpriteIdx != test.expectedTile {
			t.Errorf(fmt.Sprintf("%s failed, expected %d, got %d", test.name, test.expectedTile, res.SpriteIdx))
		}
	}
}
