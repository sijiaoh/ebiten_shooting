package units

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/sijiaoh/ebiten_shooting/time"
	"github.com/sijiaoh/ebiten_shooting/utils"
)

type PlayerActor struct {
	ActorBase

	speedPerSecond float64
}

func NewPlayerActor() PlayerActor {
	return PlayerActor{
		ActorBase:      NewActorBase(),
		speedPerSecond: 1,
	}
}

func (pa *PlayerActor) Init() {
}

func (pa *PlayerActor) Update() {
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
		speed := pa.speedPerSecond * time.Time.DeltaTime
		vec = vec.Mul(speed)
	}

	pa.pos = pa.pos.Add(vec)
}

func (pa *PlayerActor) Draw(screen *ebiten.Image) {
	screenPos := pa.pos.ToScreenPos()
	vector.DrawFilledCircle(screen, float32(screenPos.X), float32(screenPos.Y), 10, color.White, false)
}

func (pa *PlayerActor) OnDead() {
}
