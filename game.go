package main

import (
	"image/color"
	gamemap "prototype/map"
	"prototype/screen"
	"prototype/settings"
	"prototype/tiles"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	TileManager *tiles.TileManager
	Map         *gamemap.GameMap
	Screen      *screen.Screen
}

func NewGame() *Game {
	settings.Init()
	s := screen.NewScreen()
	tm := tiles.NewTileManager()
	m := gamemap.NewMap(s, tm)
	return &Game{
		TileManager: tm,
		Map:         m,
		Screen:      s,
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{
		R: 71,
		G: 45,
		B: 60,
	})
	g.Map.Render()
	g.Screen.Draw(screen)
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Screen.Width, g.Screen.Height
}
