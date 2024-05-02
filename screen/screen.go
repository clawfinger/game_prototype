package screen

import "github.com/hajimehoshi/ebiten/v2"

type Screen struct {
	Width          int
	Height         int
	Layers         [][]*spriteData
	View           ebiten.GeoM
	InternalScreen *ebiten.Image
	Op             *ebiten.DrawImageOptions
	SpriteOp       *ebiten.DrawImageOptions
}

type spriteData struct {
	sprite *ebiten.Image
	pos    ebiten.GeoM
}

func NewScreen(width int, height int) *Screen {
	return &Screen{
		Width:          width,
		Height:         height,
		Layers:         make([][]*spriteData, LayersMax),
		Op:             &ebiten.DrawImageOptions{},
		SpriteOp:       &ebiten.DrawImageOptions{},
		InternalScreen: ebiten.NewImage(width, height),
	}
}

func (s *Screen) AddToLayer(layer int, image *ebiten.Image, pos *ebiten.GeoM) {
	s.Layers[layer] = append(s.Layers[layer], &spriteData{sprite: image, pos: *pos})
}

func (s *Screen) Draw(screen *ebiten.Image) {
	s.InternalScreen.Clear()
	for _, layer := range s.Layers {
		for _, spriteData := range layer {
			s.SpriteOp.GeoM = spriteData.pos
			s.InternalScreen.DrawImage(spriteData.sprite, s.SpriteOp)
		}
	}
	screen.DrawImage(s.InternalScreen, s.Op)
	s.Layers = make([][]*spriteData, LayersMax)
}
