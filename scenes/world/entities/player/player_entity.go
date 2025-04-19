package player

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/camera"
	"github.com/sijiaoh/ebiten_shooting/game"
	"github.com/sijiaoh/ebiten_shooting/scenes/world/entities"
	"github.com/sijiaoh/ebiten_shooting/scenes/world/entities/bullets"
	"github.com/sijiaoh/ebiten_shooting/time"
)

type PlayerEntity struct {
	entities.EntityBase

	Pos            gmath.Vec
	speedPerSecond float64
}

func NewPlayerEntity() PlayerEntity {
	return PlayerEntity{
		EntityBase:     entities.NewEntityBase(),
		Pos:            gmath.Vec{},
		speedPerSecond: 2,
	}
}

func (pe *PlayerEntity) Init() {
}

func (pe *PlayerEntity) Update() {
	pe.move()
	pe.shoot()
}

func (pe *PlayerEntity) Draw(screen *ebiten.Image) {
	size := 0.5 * camera.PixelsPerUnit
	screenPos := camera.ToScreenPos(pe.Pos)
	vector.DrawFilledCircle(screen, float32(screenPos.X), float32(screenPos.Y), float32(size/2.0), color.White, false)
}

func (pe *PlayerEntity) OnDisposed() {
}

func (pe *PlayerEntity) move() {
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
		speed := pe.speedPerSecond * time.Time.DeltaTime
		vec = vec.Mulf(speed)
	}

	pe.Pos = pe.Pos.Add(vec)
}

func (pe *PlayerEntity) shoot() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		b := bullets.NewStraightBullet(pe.Pos, gmath.Vec{X: 0, Y: -1}, 5)
		game.C.Scene.AddEntity(&b)
	}
}
