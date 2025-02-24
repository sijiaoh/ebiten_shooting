package units

type Unit interface {
	IsInited() bool
	Init()
	EndInit()
	Update()
	Draw()

	IsAlive() bool
	Die()
	OnDead()
}

type UnitBase struct {
	isInited bool

	// 0以下になると死亡判定
	hp int
}

func (ub *UnitBase) NewUnitBase() {
	ub.isInited = false
	ub.hp = 1
}

func (ub *UnitBase) IsInited() bool {
	return ub.isInited
}

func (ub *UnitBase) Endinit() {
	ub.isInited = true
}

func (ub *UnitBase) IsAlive() bool {
	return ub.hp <= 0
}

func (ub *UnitBase) Die() {
	ub.hp = 0
}
