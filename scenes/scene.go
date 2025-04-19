package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sijiaoh/ebiten_shooting/entity"
)

type Scene interface {
	AddEntity(entity entity.Entity)
	Update()
	Draw(screen *ebiten.Image)
}

type SceneBase struct {
	entityManager entity.EntityManager
	drawerManager entity.DrawerManager
}

func NewSceneBase() SceneBase {
	return SceneBase{
		entityManager: entity.NewEntityManager(),
		drawerManager: entity.NewDrawerManager(),
	}
}

func (sb *SceneBase) AddEntity(entity entity.Entity) {
	sb.entityManager.AddEntity(entity)
}

func (sb *SceneBase) Update() {
	sb.entityManager.Update()
}

func (sb *SceneBase) Draw(screen *ebiten.Image) {
	sb.entityManager.Draw(&sb.drawerManager)
	sb.drawerManager.Draw(screen)
}
