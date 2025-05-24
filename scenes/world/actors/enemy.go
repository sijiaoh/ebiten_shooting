package actors

import (
	"image/color"

	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/core"
	"github.com/sijiaoh/ebiten_shooting/scenes/world/components"
)

type Enemy struct {
	ActorBase

	speedPerSecond float64
}

func NewEnemy(playerActor *PlayerActor) *Enemy {
	e := &Enemy{
		ActorBase:      *NewActorBase(),
		speedPerSecond: 1,
	}
	e.AddComponent(components.NewFollower(e, playerActor))
	return e
}

func (e *Enemy) Draw(dm *core.DrawerManager) {
	e.ActorBase.Draw(dm)

	drawer := dm.NewDrawer()
	drawer.DrawRect(e.pos, gmath.Vec{X: 0.5, Y: 0.5}, color.RGBA{255, 0, 0, 255})
	dm.AddDrawer(drawer)
}

func (e *Enemy) SpeedPerSecond() float64 {
	return e.speedPerSecond
}
