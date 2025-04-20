package entities

import (
	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/core"
)

type Entity interface {
	core.Entity

	Pos() gmath.Vec
	SetPos(pos gmath.Vec)
}

type EntityBase struct {
	core.EntityBase

	pos gmath.Vec
}

func NewEntityBase() *EntityBase {
	return &EntityBase{
		EntityBase: *core.NewEntityBase(),
	}
}

func (e *EntityBase) Pos() gmath.Vec {
	return e.pos
}

func (e *EntityBase) SetPos(pos gmath.Vec) {
	e.pos = pos
}
