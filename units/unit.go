package units

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sijiaoh/ebiten_shooting/utils"
)

type Unit interface {
	Init()
	Update()
	Draw(screen *ebiten.Image)

	OnDead()

	// 以下はUnitBaseに実装済み

	IsInited() bool
	EndInit()

	IsAlive() bool
	Die()
}

type UnitBase struct {
	isInited bool

	// 0以下になると死亡判定
	hp int

	pos utils.VectorFloat
}

func NewUnitBase() UnitBase {
	return UnitBase{
		isInited: false,
		hp:       1,
	}
}

func (ub *UnitBase) IsInited() bool {
	return ub.isInited
}

func (ub *UnitBase) EndInit() {
	ub.isInited = true
}

func (ub *UnitBase) IsAlive() bool {
	return ub.hp > 0
}

func (ub *UnitBase) Die() {
	ub.hp = 0
}
