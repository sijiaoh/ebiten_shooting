package components

import (
	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/core"
)

type FollowerEntity interface {
	core.Entity

	Pos() gmath.Vec
	SetPos(pos gmath.Vec)
	SpeedPerSecond() float64
}

type FollowerTarget interface {
	Pos() gmath.Vec
}

type Follower struct {
	core.ComponentBase
	entity FollowerEntity
	target FollowerTarget
}

func NewFollower(entity FollowerEntity, target FollowerTarget) *Follower {
	return &Follower{
		ComponentBase: *core.NewComponentBase(entity),
		entity:        entity,
		target:        target,
	}
}

func (f *Follower) Update(delta float64) {
	f.ComponentBase.Update(delta)

	targetPos := f.target.Pos()
	speed := f.entity.SpeedPerSecond()

	vector := targetPos.Sub(f.entity.Pos())
	vector = vector.Normalized()
	vector = vector.Mulf(speed * delta)

	f.entity.SetPos(f.entity.Pos().Add(vector))
}
