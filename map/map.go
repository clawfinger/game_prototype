package gamemap

import (
	gamecontext "prototype/game_context"
	"prototype/screen"
	"prototype/tiles"

	"github.com/hajimehoshi/ebiten/v2"
)

type MapSettings struct {
	SizeHorizontal int   `json:"size_horizontal"`
	SizeVertical   int   `json:"size_vertical"`
	Map            []int `json:"map"`
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
	settings := &MapSettings{
		SizeHorizontal: 10,
		SizeVertical:   10,
		Map: []int{
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
		},
	}
	m.Settings = settings
	return nil
}

func (m *GameMap) Render() {
	for row := 0; row < m.Settings.SizeVertical; row++ {
		for column := 0; column < m.Settings.SizeHorizontal; column++ {
			m.op.Reset()
			m.op.Translate(float64(column*m.tile.Settings.TileSize), float64(row*m.tile.Settings.TileSize))
			m.op.Translate(80, 40)
			m.s.AddToLayer(screen.FloorLayer, m.tile.GetTile(m.Settings.Map[m.createLinearFromRowAndColumn(row, column)]), m.op)
		}
	}
}

func (m *GameMap) createLinearFromRowAndColumn(row int, column int) int {
	return row*m.Settings.SizeHorizontal + column
}
