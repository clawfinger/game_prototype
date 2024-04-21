package gamemap

import (
	"encoding/json"
	"io"
	"os"
	"path"
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

func NewMap(s *screen.Screen, tiles *tiles.TileManager) *GameMap {
	return &GameMap{
		op:   &ebiten.GeoM{},
		s:    s,
		tile: tiles,
	}
}

func (m *GameMap) Init() error {
	settingsFile, err := os.Open(path.Join(".", "Assets", "map.json"))

	if err != nil {
		return err
	}
	byteValue, err := io.ReadAll(settingsFile)
	if err != nil {
		return err
	}
	settings := &MapSettings{}
	err = json.Unmarshal(byteValue, settings)
	if err != nil {
		return err
	}
	m.Settings = settings
	return nil
}

func (m *GameMap) Render() {
	for row := 0; row < m.Settings.SizeVertical; row++ {
		for column := 0; column < m.Settings.SizeHorizontal; column++ {
			m.op.Reset()
			m.op.Translate(float64(column*m.tile.Settings.TileSize), float64(row*m.tile.Settings.TileSize))
			m.s.AddToLayer(screen.FloorLayer, m.tile.GetTile(m.Settings.Map[m.createLinearFromRowAndColumn(row, column)]), m.op)
		}
	}
}

func (m *GameMap) createLinearFromRowAndColumn(row int, column int) int {
	return row*m.Settings.SizeHorizontal + column
}
