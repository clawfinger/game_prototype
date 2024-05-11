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
	ebiten.SetWindowTitle("Prototype")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
