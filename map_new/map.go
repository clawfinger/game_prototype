package gamemap_new

import (
	"prototype/event"
	gamecontext "prototype/game_context"
)

type Tile struct {
	SpriteIdx int
	Walkable  bool
}

type MapArena struct {
	SizeX int
	SizeY int
	Tiles []*Tile
}

type GameMap struct {
	CurrentArena *MapArena
	Arenas       []*MapArena

	ed *event.EventDispatcher
}

func NewGameMap() *GameMap {
	return &GameMap{
		Arenas: make([]*MapArena, 0),
		ed:     gamecontext.GameContext.EventDispatcher,
	}
}

func (m *GameMap) GenerateMap() {
	idxMap := []int{
		541, 542, 638, 844, 638, 638, 844, 542, 844, 844,
		844, 1, 1, 1, 1, 1, 3, 4, 1, 638,
		638, 1, 1, 1, 1, 1, 3, 3, 4, 844,
		844, 1, 1, 1, 1, 1, 1, 3, 3, 844,
		845, 1, 1, 1, 1, 1, 1, 1, 1, 445,
		638, 1, 1, 1, 1, 1, 1, 1, 1, 844,
		844, 1, 1, 1, 1, 1, 1, 1, 1, 844,
		844, 3, 1, 1, 1, 1, 1, 1, 1, 638,
		844, 3, 3, 1, 1, 1, 1, 1, 1, 638,
		844, 638, 844, 542, 844, 844, 638, 542, 844, 844,
	}

	arena := &MapArena{
		SizeX: 10,
		SizeY: 10,
	}

	var tiles []*Tile
	for i := range idxMap {
		tile := &Tile{
			SpriteIdx: idxMap[i],
		}
		tiles = append(tiles, tile)
	}
	arena.Tiles = tiles

	m.CurrentArena = arena
	m.Arenas = append(m.Arenas, arena)
	m.ed.Dispatch(&event.MapArenaLoadedEvent{})
}

func (m *GameMap) getCurrentArena() *MapArena {
	return m.CurrentArena
}
