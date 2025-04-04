package player

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/sijiaoh/ebiten_shooting/actors"
	"github.com/sijiaoh/ebiten_shooting/actors/bullets"
	"github.com/sijiaoh/ebiten_shooting/camera"
	"github.com/sijiaoh/ebiten_shooting/time"
	"github.com/sijiaoh/ebiten_shooting/utils"
)

type PlayerActor struct {
	actors.ActorBase

	Pos            utils.VectorFloat
	speedPerSecond float64
}

func NewPlayerActor() PlayerActor {
	return PlayerActor{
		ActorBase:      actors.NewActorBase(),
		Pos:            utils.VectorFloat{},
		speedPerSecond: 2,
	}
}

func (pa *PlayerActor) Init() {
}

func (pa *PlayerActor) Update() {
	pa.move()
	pa.shoot()
}

func (pa *PlayerActor) Draw(screen *ebiten.Image) {
	size := 0.5 * camera.PixelsPerUnit
	screenPos := camera.ToScreenPos(pa.Pos)
	vector.DrawFilledCircle(screen, float32(screenPos.X), float32(screenPos.Y), float32(size/2.0), color.White, false)
}

func (pa *PlayerActor) OnDead() {
}

func (pa *PlayerActor) move() {
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

	pa.Pos = pa.Pos.Add(vec)
}

func (pa *PlayerActor) shoot() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		b := bullets.NewStraightBullet(pa.Pos, utils.VectorFloat{X: 0, Y: -1}, 5)
		actors.Actors.AddActor(&b)
	}
}
