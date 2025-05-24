package actors

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/core"
	"github.com/sijiaoh/ebiten_shooting/game"
)

type PlayerActor struct {
	ActorBase

	speedPerSecond float64
}

func NewPlayerActor() *PlayerActor {
	return &PlayerActor{
		ActorBase:      *NewActorBase(),
		speedPerSecond: 2,
	}
}

func (pa *PlayerActor) Update(delta float64) {
	pa.ActorBase.Update(delta)

	pa.move(delta)
	pa.shoot()
}

func (pa *PlayerActor) Draw(dm *core.DrawerManager) {
	pa.ActorBase.Draw(dm)

	drawer := dm.NewDrawer()
	drawer.DrawCircle(pa.pos, 0.5, color.White)
	dm.AddDrawer(drawer)
}

func (pa *PlayerActor) move(delta float64) {
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
		speed := pa.speedPerSecond * delta
		vec = vec.Mulf(speed)
	}

	pa.pos = pa.pos.Add(vec)
}

func (pa *PlayerActor) shoot() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		b := NewStraightBullet(pa.pos, gmath.Vec{X: 0, Y: -1}, 5)
		game.C.Scene.AddEntity(b)
	}
}
