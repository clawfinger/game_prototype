package screen

import "github.com/hajimehoshi/ebiten/v2"

type Screen struct {
	Width          int
	Height         int
	LayersNum      int
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

func NewScreen(layersNums int, width int, height int) *Screen {
	return &Screen{
		Width:          width,
		Height:         height,
		LayersNum:      layersNums,
		Layers:         make([][]*spriteData, layersNums),
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
	for i, layer := range s.Layers {
		for _, spriteData := range layer {
			s.SpriteOp.GeoM = spriteData.pos
			s.InternalScreen.DrawImage(spriteData.sprite, s.SpriteOp)
		}
		s.Layers[i] = s.Layers[i][:1]
	}
	screen.DrawImage(s.InternalScreen, s.Op)
}
