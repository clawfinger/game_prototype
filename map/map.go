package gamemap

import (
	gamecontext "prototype/game_context"
	"prototype/screen"
	"prototype/settings"
	"prototype/tiles"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	SpriteIdx int
	X         int
	Y         int
}

type MapSettings struct {
	SizeHorizontal int     `json:"size_horizontal"`
	SizeVertical   int     `json:"size_vertical"`
	Map            []*Tile `json:"map"`
}

type GameMap struct {
	Settings *MapSettings
	op       *ebiten.GeoM
	s        *screen.Screen
	tile     *tiles.TileManager
}

func NewMap() *GameMap {
	return &GameMap{
		op:   &ebiten.GeoM{},
		s:    gamecontext.GameContext.Screen,
		tile: gamecontext.GameContext.TileManager,
	}
}

func (m *GameMap) Init() error {
	idxMap := []int{
		541, 542, 638, 844, 638, 638, 844, 542, 844, 844,
		844, 1, 1, 1, 1, 1, 3, 4, 1, 638,
		638, 1, 1, 1, 1, 1, 3, 3, 4, 844,
		844, 1, 1, 1, 1, 1, 1, 3, 3, 844,
		845, 1, 1, 1, 1, 1, 1, 1, 1, 445,
		638, 1, 1, 1, 1, 1, 1, 1, 1, 844,
		844, 1, 1, 1, 1, 1, 1, 1, 1, 844,
		844, 3, 1, 1, 1, 1, 1, 1, 1, 638,
		844, 3, 3, 1, 1, 1, 1, 1, 1, 638,
		844, 638, 844, 542, 844, 844, 638, 542, 844, 844,
	}
	settings := &MapSettings{
		SizeHorizontal: 10,
		SizeVertical:   10,
	}
	m.Settings = settings
	var tiles []*Tile
	for i := range idxMap {
		x, y := m.createRowAndColumnFromLiear(i)
		tile := &Tile{
			SpriteIdx: idxMap[i],
			X:         x,
			Y:         y,
		}
		tiles = append(tiles, tile)
	}
	settings.Map = tiles

	return nil
}

func (m *GameMap) GetLevelSizeInPixels() (int, int) {
	return m.Settings.SizeHorizontal * settings.Data.Tilesheet.TileSize, m.Settings.SizeVertical * settings.Data.Tilesheet.TileSize
}

func (m *GameMap) createLinearFromRowAndColumn(row int, column int) int {
	return row*m.Settings.SizeHorizontal + column
}

func (m *GameMap) createRowAndColumnFromLiear(idx int) (int, int) {
	x := idx % m.Settings.SizeHorizontal
	y := idx / m.Settings.SizeHorizontal
	return x, y
}

func (m *GameMap) GetTile(x, y int) (*Tile, bool) {
	within := m.isWithinMap(x, y)
	if !within {
		return nil, false
	}
	tileX := x / settings.Data.Tilesheet.TileSize
	tileY := y / settings.Data.Tilesheet.TileSize
	return m.Settings.Map[m.createLinearFromRowAndColumn(tileX, tileY)], true
}

func (m *GameMap) isWithinMap(x, y int) bool {
	sizeX, sizeY := m.GetLevelSizeInPixels()
	if x >= 0 && x <= sizeX &&
		y >= 0 && y <= sizeY {
		return true
	}
	return false
}
