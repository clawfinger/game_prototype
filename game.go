package main

import (
	"image/color"
	"prototype/game"
	gamecontext "prototype/game_context"
	"prototype/settings"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	GameInstance *game.GameInstance
}

func NewGame() (*Game, error) {
	err := settings.Init()
	if err != nil {
		return nil, err
	}
	gamecontext.Init()

	gi := game.NewGameInstance()
	gi.Init()
	return &Game{
		GameInstance: gi,
	}, nil
}

func (g *Game) Update() error {
	g.GameInstance.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{
		R: 71,
		G: 45,
		B: 60,
	})
	g.GameInstance.Render()
	gamecontext.GameContext.Screen.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gamecontext.GameContext.Screen.Width, gamecontext.GameContext.Screen.Height
}
