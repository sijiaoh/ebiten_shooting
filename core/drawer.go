package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/camera"
)

// 意図した順番で描画を行うための仕組み
// EntityのDraw内では直接描画を行わず、drawerの生成を行う
// drawerは指定したLayerやOrderに基づいてソートされてから描画する
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

func (d *drawer) DrawCircle(pos gmath.Vec, radius float64, clr color.Color) {
	d.Draw = func(screen *ebiten.Image) {
		size := radius * camera.PixelsPerUnit
		screenPos := camera.ToScreenPos(pos)
		vector.DrawFilledCircle(screen, float32(screenPos.X), float32(screenPos.Y), float32(size/2.0), clr, false)
	}
}

func (d *drawer) DrawRect(pos gmath.Vec, size gmath.Vec, clr color.Color) {
	d.Draw = func(screen *ebiten.Image) {
		size = size.Mulf(camera.PixelsPerUnit)
		screenPos := camera.ToScreenPos(pos)
		leftTop := screenPos.Add(gmath.Vec{X: -size.X / 2, Y: -size.Y / 2})
		vector.DrawFilledRect(screen, float32(leftTop.X), float32(leftTop.Y), float32(size.X), float32(size.Y), clr, false)
	}
}
