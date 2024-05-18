package tiles

import (
	"image"
	"path"
	"prototype/settings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TileManager struct {
	TileSheet *ebiten.Image
}

func NewTileManager() *TileManager {
	return &TileManager{}
}

func (m *TileManager) Init() error {
	tilemap, _, err := ebitenutil.NewImageFromFile(path.Join(".", "Assets", "Tilesheet", "colored-transparent.png"))

	if err != nil {
		return err
	}
	m.TileSheet = tilemap
	return nil
}

func (m *TileManager) GetTile(number int) *ebiten.Image {
	if number != 0 {
		number -= 1
	}
	row := number / settings.Data.Tilesheet.TotalHorizontal
	column := number % settings.Data.Tilesheet.TotalHorizontal
	x := column*settings.Data.Tilesheet.TileSize + (column-1)*settings.Data.Tilesheet.TileOffset
	y := row*settings.Data.Tilesheet.TileSize + (row-1)*settings.Data.Tilesheet.TileOffset
	return m.TileSheet.SubImage(image.Rect(x, y, settings.Data.Tilesheet.TileSize+x, settings.Data.Tilesheet.TileSize+y)).(*ebiten.Image)
}
