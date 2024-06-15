package gamemap

import (
	"prototype/event"
	gamecontext "prototype/game_context"
	"prototype/screen"
	"prototype/settings"
	"prototype/tiles"

	"github.com/hajimehoshi/ebiten/v2"
)

type MapRenderer struct {
	m             *GameMap
	mapTileOp     *ebiten.GeoM
	mapPositionOp *ebiten.GeoM
	tile          *tiles.TileManager
	screen        *screen.Screen
	ed            *event.EventDispatcher
	image         *ebiten.Image
	currentArena  *MapArena
}

func NewMapRenderer(m *GameMap) *MapRenderer {
	mapPos := &ebiten.GeoM{}
	mapPos.Translate(float64(settings.GameplayData.MapOffsetX), float64(settings.GameplayData.MapOffsetY))
	return &MapRenderer{
		m:             m,
		mapTileOp:     &ebiten.GeoM{},
		tile:          gamecontext.GameContext.TileManager,
		screen:        gamecontext.GameContext.Screen,
		mapPositionOp: mapPos,
		ed:            gamecontext.GameContext.EventDispatcher,
	}
}

func (r *MapRenderer) Init() {
	r.ed.Subscribe(event.MapArenaLoadedEventName, r)
	r.ed.Subscribe(event.RenderEventName, r)
}

func (r *MapRenderer) Render() {
	r.screen.AddToLayer(screen.FloorLayer, r.image, r.mapPositionOp)
}

func (r *MapRenderer) Notify(e event.Event) {
	switch event := e.(type) {
	case *event.MapArenaLoadedEvent:
		r.handleMapArenaLoadedEvent(event)
	case *event.RenderEvent:
		r.Render()
	}
}

func createRowAndColumnFromLiear(idx int, sizeX int) (int, int) {
	x := idx % sizeX
	y := idx / sizeX
	return x, y
}

func createLinearFromRowAndColumn(row int, column int, sizeX int) int {
	return row*sizeX + column
}

func (r *MapRenderer) handleMapArenaLoadedEvent(e *event.MapArenaLoadedEvent) {
	r.currentArena = r.m.getCurrentArena()

	width := settings.Data.Tilesheet.TileSize * r.currentArena.SizeX
	height := settings.Data.Tilesheet.TileSize * r.currentArena.SizeY
	r.image = ebiten.NewImage(width, height)

	for i := range r.currentArena.Tiles {
		tile := r.currentArena.Tiles[i]
		x, y := createRowAndColumnFromLiear(i, r.currentArena.SizeX)

		r.mapTileOp.Reset()
		r.mapTileOp.Translate(float64(x*settings.Data.Tilesheet.TileSize), float64(y*settings.Data.Tilesheet.TileSize))
		r.image.DrawImage(r.tile.GetTile(tile.SpriteIdx), &ebiten.DrawImageOptions{
			GeoM: *r.mapTileOp,
		})
	}
}
