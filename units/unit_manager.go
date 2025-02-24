package units

type UnitManager struct {
	units []Unit
}

// OnDeadを確実に呼び出すためにもRemoveUnitは提供しない
func (um *UnitManager) AddUnit(unit Unit) {
	um.units = append(um.units, unit)
}

func (um *UnitManager) Update(unit Unit) {
	um.initUnits()

	for _, unit := range um.units {
		if unit.IsAlive() {
			unit.Update()
		}
	}

	um.removeDeadUnits()
}

func (um *UnitManager) Draw(unit Unit) {
	for _, unit := range um.units {
		unit.Draw()
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
	var aliveUnits []Unit
	for _, unit := range um.units {
		if unit.IsAlive() {
			aliveUnits = append(aliveUnits, unit)
		} else {
			unit.OnDead()
		}
	}
	um.units = aliveUnits
}
