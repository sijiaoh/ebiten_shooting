package units

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/sijiaoh/ebiten_shooting/time"
	"github.com/sijiaoh/ebiten_shooting/utils"
)

type PlayerUnit struct {
	UnitBase

	speedPerSecond float64
}

func NewPlayerUnit() PlayerUnit {
	return PlayerUnit{
		UnitBase:       NewUnitBase(),
		speedPerSecond: 1,
	}
}

func (pu *PlayerUnit) Init() {
}

func (pu *PlayerUnit) Update() {
	vec := utils.VectorFloat{}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		vec.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		vec.Y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		vec.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		vec.X += 1
	}

	if vec.LengthSquared() > 0 {
		vec = vec.Normalize()
		speed := pu.speedPerSecond * time.Time.DeltaTime
		vec = vec.Mul(speed)
	}

	pu.pos = pu.pos.Add(vec)
}

func (pu *PlayerUnit) Draw(screen *ebiten.Image) {
	screenPos := pu.pos.ToScreenPos()
	vector.DrawFilledCircle(screen, float32(screenPos.X), float32(screenPos.Y), 10, color.White, false)
}

func (pu *PlayerUnit) OnDead() {
}
