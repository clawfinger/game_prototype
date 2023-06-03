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

	"github.com/hajimehoshi/ebiten/v2"
)

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
