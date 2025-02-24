package units

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerUnit struct {
	UnitBase
}

func (pu *PlayerUnit) NewPlayerUnit() {
	pu.NewUnitBase()
	pu.hp = 3
}

func (pu *PlayerUnit) Init() {
}

func (pu *PlayerUnit) Update() {
}

func (pu *PlayerUnit) Draw(screen *ebiten.Image) {
}

func (pu *PlayerUnit) OnDead() {
}
