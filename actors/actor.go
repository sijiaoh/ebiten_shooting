package actors

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Actor interface {
	Init()
	Update()
	Draw(screen *ebiten.Image)

	OnDead()

	// 以下はActorBaseに実装済み

	IsInited() bool
	EndInit()
	IsAlive() bool
}

type ActorBase struct {
	isInited bool
}

func NewActorBase() ActorBase {
	return ActorBase{
		isInited: false,
	}
}

func (ub *ActorBase) IsInited() bool {
	return ub.isInited
}

func (ub *ActorBase) EndInit() {
	ub.isInited = true
}

func (ub *ActorBase) IsAlive() bool {
	return true
}
