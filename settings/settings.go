package settings

var Settings GameSettings

type GameSettings struct {
	ScreenSizeX int `json:"screenSizeX"`
	ScreenSizeY int `json:"screenSizeY"`
}

func Init() {
	Settings = GameSettings{
		ScreenSizeX: 320,
		ScreenSizeY: 240,
	}
}
