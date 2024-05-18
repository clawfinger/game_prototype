package gamemap

import (
	"prototype/event"
	gamecontext "prototype/game_context"
	"prototype/screen"
	"prototype/settings"
	"prototype/tiles"

	"github.com/hajimehoshi/ebiten/v2"
)

type CursorHandler struct {
	ed     *event.EventDispatcher
	screen *screen.Screen
}

func NewCursorHandler() *CursorHandler {
	return &CursorHandler{
		ed:     gamecontext.GameContext.EventDispatcher,
		screen: gamecontext.GameContext.Screen,
	}
}

func (c *CursorHandler) Init() {
	c.ed.Subscribe(event.MouseMovedEventName, c)
}

func (c *CursorHandler) Update() {

}

func (c *CursorHandler) Notify(e event.Event) {
	switch event := e.(type) {
	case *event.MouseMovedEvent:
		c.handleMouseMovedEvent(event)
	}
}

func (c *CursorHandler) handleMouseMovedEvent(e *event.MouseMovedEvent) {

}

type MapRenderer struct {
	level *GameMap
	ch    *CursorHandler
	image *ebiten.Image

	op            *ebiten.GeoM
	mapPositionOp *ebiten.GeoM
	tile          *tiles.TileManager
	screen        *screen.Screen
}

func NewMapRenderer(m *GameMap) *MapRenderer {
	mapPos := &ebiten.GeoM{}
	mapPos.Translate(float64(settings.GameplayData.MapOffsetX), float64(settings.GameplayData.MapOffsetY))
	return &MapRenderer{
		level:         m,
		ch:            NewCursorHandler(),
		op:            &ebiten.GeoM{},
		tile:          gamecontext.GameContext.TileManager,
		screen:        gamecontext.GameContext.Screen,
		mapPositionOp: mapPos,
	}
}

func (r *MapRenderer) Init() {
	sizeX, sizeY := r.level.GetLevelSizeInPixels()
	levelImage := ebiten.NewImage(sizeX, sizeY)
	r.image = levelImage
}

func (r *MapRenderer) Render() {
	for row := 0; row < r.level.Settings.SizeVertical; row++ {
		for column := 0; column < r.level.Settings.SizeHorizontal; column++ {
			r.op.Reset()
			r.op.Translate(float64(column*settings.Data.Tilesheet.TileSize), float64(row*settings.Data.Tilesheet.TileSize))
			r.op.Concat(*r.mapPositionOp)
			r.screen.AddToLayer(screen.FloorLayer, r.tile.GetTile(r.level.Settings.Map[r.level.createLinearFromRowAndColumn(row, column)].SpriteIdx), r.op)
		}
	}
}
