package input

import (
	"math"
	"prototype/event"

	"github.com/hajimehoshi/ebiten/v2"
)

type InputState int

const (
	LeftButtonPressed = iota
	LeftButtonReleased
	// RightButtonPressed
	// RightButtonReleased
)

type InputController struct {
	mouseState InputState
	ed         *event.EventDispatcher
	LastMouseX int
	LastMouseY int
}

func NewInputController(ed *event.EventDispatcher) *InputController {
	return &InputController{
		ed:         ed,
		mouseState: LeftButtonReleased,
	}
}

func (c *InputController) Update() {
	cursorX, cursorY := ebiten.CursorPosition()

	if math.Abs(float64(cursorX)-float64(c.LastMouseX)) > 0.1 ||
		math.Abs(float64(cursorY)-float64(c.LastMouseY)) > 0.1 {
		c.ed.Dispatch(&event.MouseMovedEvent{
			X: cursorX,
			Y: cursorY,
		})
		c.LastMouseX = cursorX
		c.LastMouseY = cursorY
	}

	switch c.mouseState {
	case LeftButtonReleased:
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			c.ed.Dispatch(&event.MousePressedEvent{
				X: cursorX,
				Y: cursorY,
			})
			c.mouseState = LeftButtonPressed
		}
	case LeftButtonPressed:
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			c.ed.Dispatch(&event.MouseRelesedEvent{
				X: cursorX,
				Y: cursorY,
			})
			c.mouseState = LeftButtonReleased
		}
	}
}
