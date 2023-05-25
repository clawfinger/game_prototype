package main

// Tile size                 •  16px × 16px
// Space between tiles       •  1px × 1px
// ---
// Total tiles (horizontal)  •  49 tiles
// Total tiles (vertical)    •  22 tiles
// ---
// Total tiles in sheet      •  1078 tiles

import (
	"log"
	gamemap "prototype/map"
	"prototype/tiles"

	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	TileManager *tiles.TileManager
	Map         *gamemap.GameMap
}

func NewGame() *Game {
	tm := tiles.NewTileManager()
	m := gamemap.NewMap()
	return &Game{
		TileManager: tm,
		Map:         m,
	}
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	screen.Fill(color.RGBA{
		R: 71,
		G: 45,
		B: 60,
	})
	g.Map.Draw(screen, g.TileManager)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Prototype")
	game := NewGame()
	if err := game.TileManager.Init(); err != nil {
		log.Fatal(err)
	}
	if err := game.Map.Init(); err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
