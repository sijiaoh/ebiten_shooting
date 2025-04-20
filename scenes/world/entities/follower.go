package entities

import (
	"github.com/sijiaoh/ebiten_shooting/core"
)

type Follower struct {
	core.ComponentBase
	target Entity
}

func NewFollower(entity Entity, target Entity) *Follower {
	return &Follower{
		ComponentBase: *core.NewComponentBase(entity),
		target:        target,
	}
}

func (f *Follower) Awake() {
}

func (f *Follower) Init() {
}

func (f *Follower) Update(delta float64) {
}

func (f *Follower) Draw(dm *core.DrawerManager) {
}

func (f *Follower) OnDisposed() {
}
