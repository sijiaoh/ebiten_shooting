package core

import "github.com/hajimehoshi/ebiten/v2"

type Drawer struct {
	Layer int
	Order float64
	Draw  func(screen *ebiten.Image)
}
