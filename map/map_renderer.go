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
	m     *GameMap
	image *ebiten.Image

	cursorOp      *ebiten.GeoM
	mapTileOp     *ebiten.GeoM
	mapPositionOp *ebiten.GeoM
	tile          *tiles.TileManager
	screen        *screen.Screen
	ed            *event.EventDispatcher
}

func NewMapRenderer(m *GameMap) *MapRenderer {
	mapPos := &ebiten.GeoM{}
	mapPos.Translate(float64(settings.GameplayData.MapOffsetX), float64(settings.GameplayData.MapOffsetY))
	return &MapRenderer{
		m:             m,
		mapTileOp:     &ebiten.GeoM{},
		cursorOp:      &ebiten.GeoM{},
		tile:          gamecontext.GameContext.TileManager,
		screen:        gamecontext.GameContext.Screen,
		mapPositionOp: mapPos,
		ed:            gamecontext.GameContext.EventDispatcher,
	}
}

func (r *MapRenderer) Init() {
	sizeX, sizeY := r.m.GetLevelSizeInPixels()
	levelImage := ebiten.NewImage(sizeX, sizeY)
	r.image = levelImage
	// r.ed.Subscribe(event.MouseMovedEventName, r)
	r.ed.Subscribe(event.RenderEventName, r)

}

func (r *MapRenderer) Render() {
	for row := 0; row < r.m.Settings.SizeVertical; row++ {
		for column := 0; column < r.m.Settings.SizeHorizontal; column++ {
			r.mapTileOp.Reset()
			r.mapTileOp.Translate(float64(column*settings.Data.Tilesheet.TileSize), float64(row*settings.Data.Tilesheet.TileSize))
			r.mapTileOp.Concat(*r.mapPositionOp)
			r.screen.AddToLayer(screen.FloorLayer, r.tile.GetTile(r.m.Settings.Map[r.m.createLinearFromRowAndColumn(row, column)].SpriteIdx), r.mapTileOp)
		}
	}
	r.cursorOp.Reset()
	mouseX, mouseY := ebiten.CursorPosition()
	op := *r.mapPositionOp
	op.Invert()
	mapX, mapY := op.Apply(float64(mouseX), float64(mouseY))
	tile, found := r.m.GetTile(int(mapX), int(mapY))
	if found {
		x := tile.X*settings.Data.Tilesheet.TileSize + (tile.X-1)*settings.Data.Tilesheet.TileOffset
		y := tile.Y*settings.Data.Tilesheet.TileSize + (tile.Y-1)*settings.Data.Tilesheet.TileOffset
		r.cursorOp.Translate(float64(x), float64(y))
		r.screen.AddToLayer(screen.GuiLayer, r.tile.GetTile(625), r.cursorOp)
	}

}

func (r *MapRenderer) Notify(e event.Event) {
	switch event := e.(type) {
	case *event.MouseMovedEvent:
		r.handleMouseMovedEvent(event)
	case *event.RenderEvent:
		r.Render()
	}
}

func (r *MapRenderer) handleMouseMovedEvent(e *event.MouseMovedEvent) {
	r.cursorOp.Reset()
	r.cursorOp.Translate(float64(e.X), float64(e.Y))
}
