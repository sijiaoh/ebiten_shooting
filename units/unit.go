package units

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sijiaoh/ebiten_shooting/utils"
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
	Die()
}

type ActorBase struct {
	isInited bool

	// 0以下になると死亡判定
	hp int

	pos utils.VectorFloat
}

func NewActorBase() ActorBase {
	return ActorBase{
		isInited: false,
		hp:       1,
	}
}

func (ub *ActorBase) IsInited() bool {
	return ub.isInited
}

func (ub *ActorBase) EndInit() {
	ub.isInited = true
}

func (ub *ActorBase) IsAlive() bool {
	return ub.hp > 0
}

func (ub *ActorBase) Die() {
	ub.hp = 0
}
