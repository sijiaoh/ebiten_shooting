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
	entityManager *entityManager
	drawerManager *DrawerManager
}

func NewSceneBase() SceneBase {
	return SceneBase{
		entityManager: newEntityManager(),
		drawerManager: newDrawerManager(),
	}
}

func (sb *SceneBase) AddEntity(entity Entity) {
	sb.entityManager.addEntity(entity)
}

func (sb *SceneBase) Update() {
	sb.entityManager.update()
}

func (sb *SceneBase) Draw(screen *ebiten.Image) {
	sb.entityManager.draw(sb.drawerManager)
	sb.drawerManager.draw(screen)
	sb.drawerManager.clearDrawers()
}
