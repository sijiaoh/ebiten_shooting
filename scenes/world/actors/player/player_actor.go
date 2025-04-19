package player

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/camera"
	"github.com/sijiaoh/ebiten_shooting/game"
	"github.com/sijiaoh/ebiten_shooting/scenes/world/actors"
	"github.com/sijiaoh/ebiten_shooting/scenes/world/actors/bullets"
	"github.com/sijiaoh/ebiten_shooting/time"
)

type PlayerActor struct {
	actors.ActorBase

	Pos            gmath.Vec
	speedPerSecond float64
}

func NewPlayerActor() PlayerActor {
	return PlayerActor{
		ActorBase:      actors.NewActorBase(),
		Pos:            gmath.Vec{},
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

func (pa *PlayerActor) OnDisposed() {
}

func (pa *PlayerActor) move() {
	vec := gmath.Vec{}
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

	if vec.LenSquared() > 0 {
		vec = vec.Normalized()
		speed := pa.speedPerSecond * time.Time.DeltaTime
		vec = vec.Mulf(speed)
	}

	pa.Pos = pa.Pos.Add(vec)
}

func (pa *PlayerActor) shoot() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		b := bullets.NewStraightBullet(pa.Pos, gmath.Vec{X: 0, Y: -1}, 5)
		game.C.Scene.AddActor(&b)
	}
}
