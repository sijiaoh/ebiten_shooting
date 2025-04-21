package core

import "github.com/hajimehoshi/ebiten/v2"

type drawer struct {
	Layer int
	Order float64
	Draw  func(screen *ebiten.Image)
}

func (d *drawer) reset() {
	d.Layer = 0
	d.Order = 0
	d.Draw = nil
}
