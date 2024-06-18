package gamemap

import (
	"fmt"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestLinear(t *testing.T) {
	m := GameMap{
		CurrentArena: &MapArena{
			SizeX: 5,
			SizeY: 7,
		},
	}
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
	m.CurrentArena.Tiles = tiles
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
		if res := createLinearFromRowAndColumn(test.row, test.column, m.CurrentArena.SizeX); m.CurrentArena.Tiles[res].SpriteIdx != test.expected {
			t.Errorf(fmt.Sprintf("%s failed, expected %d, got %d, index %d", test.name, test.expected, m.CurrentArena.Tiles[res].SpriteIdx, res))
		}
	}
}

func TestFromLinear(t *testing.T) {
	m := GameMap{
		CurrentArena: &MapArena{
			SizeX: 5,
			SizeY: 7,
		},
	}
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
	m.CurrentArena.Tiles = tiles
	tests := []struct {
		name           string
		expectedRow    int
		expectedColumn int
		linear         int
	}{
		{
			name:           "simple",
			linear:         6,
			expectedRow:    1,
			expectedColumn: 1,
		},
		{
			name:           "0",
			linear:         0,
			expectedRow:    0,
			expectedColumn: 0,
		},
		{
			name:           "",
			linear:         4,
			expectedRow:    0,
			expectedColumn: 4,
		},
	}
	for _, test := range tests {
		x, y := createRowAndColumnFromLiear(test.linear, m.CurrentArena.SizeX)
		if x != test.expectedColumn {
			t.Errorf(fmt.Sprintf("%s failed, expected column %d, got %d, linear %d", test.name, test.expectedColumn, x, test.linear))
		}
		if y != test.expectedRow {
			t.Errorf(fmt.Sprintf("%s failed, expected row %d, got %d, linear %d", test.name, test.expectedRow, y, test.linear))
		}
	}
}

// func TestGetTiles(t *testing.T) {
// 	m := GameMap{
// 		CurrentArena: &MapArena{
// 			SizeX: 5,
// 			SizeY: 7,
// 		},
// 	}
// 	idxMap := []int{
// 		7, 1, 1, 1, 1,
// 		1, 5, 2, 2, 1,
// 		1, 2, 2, 2, 1,
// 		1, 2, 2, 2, 1,
// 		1, 2, 2, 2, 1,
// 		1, 2, 2, 2, 1,
// 		1, 1, 1, 1, 9,
// 	}
// 	var tiles []*Tile
// 	for i := range idxMap {
// 		tiles = append(tiles, &Tile{SpriteIdx: idxMap[i]})
// 	}
// 	m.CurrentArena.Tiles = tiles
// 	tests := []struct {
// 		name          string
// 		clickX        int
// 		clickY        int
// 		expectedTile  int
// 		expectedFound bool
// 	}{
// 		{
// 			name:          "not found",
// 			clickX:        10000,
// 			clickY:        10000,
// 			expectedFound: false,
// 		},
// 		{
// 			name:          "found",
// 			clickX:        16 + 8,
// 			clickY:        16 + 8,
// 			expectedTile:  99,
// 			expectedFound: true,
// 		},
// 		{
// 			name:          "found between",
// 			clickX:        16,
// 			clickY:        16,
// 			expectedTile:  99,
// 			expectedFound: true,
// 		},
// 	}
// 	for _, test := range tests {
// 		res, ok := m.GetTile(test.clickX, test.clickY)
// 		if !test.expectedFound {
// 			if ok != test.expectedFound {
// 				t.Errorf("not found not es axpected")
// 			}
// 			continue
// 		}

// 		if res.SpriteIdx != test.expectedTile {
// 			t.Errorf(fmt.Sprintf("%s failed, expected %d, got %d", test.name, test.expectedTile, res.SpriteIdx))
// 		}
// 	}
// }

func TestCoordsTranslate(t *testing.T) {
	translate := ebiten.GeoM{}
	translate.Reset()

	translate.Translate(10, 10)

	translate.Invert()
	resX, resY := translate.Apply(5, 5)
	t.Logf("%f:%f\n", resX, resY)
}
