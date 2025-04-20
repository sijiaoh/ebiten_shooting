package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	AddEntity(entity Entity)
	Update()
	Draw(screen *ebiten.Image)
}

type SceneBase struct {
	entityManager *EntityManager
	drawerManager *DrawerManager
}

func NewSceneBase() SceneBase {
	return SceneBase{
		entityManager: NewEntityManager(),
		drawerManager: NewDrawerManager(),
	}
}

func (sb *SceneBase) AddEntity(entity Entity) {
	sb.entityManager.AddEntity(entity)
}

func (sb *SceneBase) Update() {
	sb.entityManager.Update()
}

func (sb *SceneBase) Draw(screen *ebiten.Image) {
	sb.entityManager.Draw(sb.drawerManager)
	sb.drawerManager.Draw(screen)
	sb.drawerManager.ClearDrawers()
}
