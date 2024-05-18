package settings

import (
	"encoding/json"
	"io"
	"os"
	"path"
)

var Data GameSettings

// 468 - character, 275 - enemy // 625 - target mark
type TileSheetSettings struct {
	TileSize        int `json:"tile_size"`
	TileOffset      int `json:"tile_offset"`
	TotalHorizontal int `json:"total_horizontal"`
	TotalVertical   int `json:"total_vertical"`
}

type ScreenSettings struct {
	ScreenSizeX int `json:"screenSizeX"`
	ScreenSizeY int `json:"screenSizeY"`
}

type GameSettings struct {
	Screen    ScreenSettings    `json:"screen"`
	Tilesheet TileSheetSettings `json:"tilesheet"`
}

var GameplayData GameplaySettings

type GameplaySettings struct {
	MapOffsetX int `json:"mapOffsetX"`
	MapOffsetY int `json:"mapOffsetY"`
}

func Init() error {
	Data = GameSettings{}
	Data.load()

	GameplayData = GameplaySettings{
		MapOffsetX: 80,
		MapOffsetY: 40,
	}
	return nil
}

func (s *GameSettings) load() error {
	settingsFile, err := os.Open(path.Join(".", "Assets", "settings.json"))
	if err != nil {
		return err
	}
	rawFile, err := io.ReadAll(settingsFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawFile, s)
	if err != nil {
		return err
	}
	return nil
}
