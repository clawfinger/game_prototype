package tiles

import (
	"encoding/json"
	"image"
	"io"
	"os"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// 468 - character, 275 - enemy // 625 - target mark
type TileSheetSettings struct {
	TileSize        int `json:"tile_size"`
	TileOffset      int `json:"tile_offset"`
	TotalHorizontal int `json:"total_horizontal"`
	TotalVertical   int `json:"total_vertical"`
}

type TileManager struct {
	TileSheet *ebiten.Image
	Settings  *TileSheetSettings
}

func NewTileManager() *TileManager {
	return &TileManager{}
}

func (m *TileManager) Init() error {
	tilemap, _, err := ebitenutil.NewImageFromFile(path.Join(".", "Assets", "Tilesheet", "colored-transparent.png"))

	if err != nil {
		return err
	}
	settingsFile, err := os.Open(path.Join(".", "Assets", "Tilesheet", "tilesheet.json"))
	if err != nil {
		return err
	}
	byteValue, err := io.ReadAll(settingsFile)
	if err != nil {
		return err
	}
	settings := &TileSheetSettings{}
	err = json.Unmarshal(byteValue, settings)
	if err != nil {
		return err
	}
	m.TileSheet = tilemap
	m.Settings = settings
	return nil
}

func (m *TileManager) GetTile(number int) *ebiten.Image {
	if number != 0 {
		number -= 1
	}
	row := number / m.Settings.TotalHorizontal
	column := number % m.Settings.TotalHorizontal
	x := column*m.Settings.TileSize + (column-1)*m.Settings.TileOffset
	y := row*m.Settings.TileSize + (row-1)*m.Settings.TileOffset
	return m.TileSheet.SubImage(image.Rect(x, y, m.Settings.TileSize+x, m.Settings.TileSize+y)).(*ebiten.Image)
}
