package gamemap_new

import (
	"prototype/event"
	gamecontext "prototype/game_context"
	"prototype/screen"
	"prototype/settings"
	"prototype/tiles"

	"github.com/hajimehoshi/ebiten"
)

type MapRenderer struct {
	m             *GameMap
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
		tile:          gamecontext.GameContext.TileManager,
		screen:        gamecontext.GameContext.Screen,
		mapPositionOp: mapPos,
		ed:            gamecontext.GameContext.EventDispatcher,
	}
}

func (r *MapRenderer) Init() {
	r.ed.Subscribe(event.MapArenaLoadedEventName, r)
}

func (r *MapRenderer) Render() {

}

func (r *MapRenderer) Notify(e event.Event) {
	switch event := e.(type) {
	case *event.MapArenaLoadedEvent:
		r.handleMapArenaLoadedEvent(event)
	}
}

func (r *MapRenderer) handleMapArenaLoadedEvent(e *event.MapArenaLoadedEvent) {
	// arena := r.m.getCurrentArena()
}
