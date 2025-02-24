package units

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type PlayerUnit struct {
	UnitBase
}

func (pu *PlayerUnit) NewPlayerUnit() {
	pu.NewUnitBase()
	pu.hp = 3
}

func (pu *PlayerUnit) Init() {
}

func (pu *PlayerUnit) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		pu.pos.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		pu.pos.Y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		pu.pos.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		pu.pos.X += 1
	}
}

func (pu *PlayerUnit) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(pu.pos.X), float32(pu.pos.Y), 10, color.White, false)
}

func (pu *PlayerUnit) OnDead() {
}
