package gui

import (
	"prototype/event"
	gamecontext "prototype/game_context"
	"prototype/screen"
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
