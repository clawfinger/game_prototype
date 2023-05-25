package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MouseState struct {
	Pressed bool
	Button  ebiten.MouseButton
}

type InputController struct {
	MouseState MouseState
}

func (c *InputController) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	}
	// 	c.MouseState.Pressed = true
	// 	c.MouseState.Button = ebiten.MouseButtonLeft
	// 	// fire left mouse button pressed event
	// } else {
	// 	if c.MouseState.Pressed
	// }
}
