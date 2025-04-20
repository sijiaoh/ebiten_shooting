package actors

import (
	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/core"
)

type Actor interface {
	core.Entity

	Pos() gmath.Vec
	SetPos(pos gmath.Vec)
}

type ActorBase struct {
	core.EntityBase

	pos gmath.Vec
}

func NewActorBase() *ActorBase {
	return &ActorBase{
		EntityBase: *core.NewEntityBase(),
	}
}

func (e *ActorBase) Pos() gmath.Vec {
	return e.pos
}

func (e *ActorBase) SetPos(pos gmath.Vec) {
	e.pos = pos
}
