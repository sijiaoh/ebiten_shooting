package manager

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sijiaoh/ebiten_shooting/units"
)

type UnitManager struct {
	units []units.Unit
}

var um = &UnitManager{}

// OnDeadを確実に呼び出すためにもRemoveUnitは提供しない
func AddUnit(unit units.Unit) {
	um.units = append(um.units, unit)
}

func Update() {
	um.initUnits()

	for _, unit := range um.units {
		if unit.IsAlive() {
			unit.Update()
		}
	}

	um.removeDeadUnits()
}

func Draw(screen *ebiten.Image) {
	for _, unit := range um.units {
		unit.Draw(screen)
	}
}

func (um *UnitManager) initUnits() {
	for _, unit := range um.units {
		if unit.IsAlive() && !unit.IsInited() {
			unit.Init()
			unit.EndInit()
		}
	}
}

func (um *UnitManager) removeDeadUnits() {
	var aliveUnits []units.Unit
	for _, unit := range um.units {
		if unit.IsAlive() {
			aliveUnits = append(aliveUnits, unit)
		} else {
			unit.OnDead()
		}
	}
	um.units = aliveUnits
}
