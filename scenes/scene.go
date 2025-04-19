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
	entities      []entity.Entity
	addedEntities []entity.Entity
}

// OnDisposedを確実に呼び出すためにもRemoveEntityは提供しない
func (sb *SceneBase) AddEntity(entity entity.Entity) {
	sb.addedEntities = append(sb.addedEntities, entity)
}

func (sb *SceneBase) Update() {
	sb.initEntities()

	for _, entity := range sb.entities {
		if entity.IsActive() {
			entity.Update(1.0 / 60 * entity.TimeScale())
		}
	}

	sb.removeDisposedEntities()
}

func (sb *SceneBase) Draw(screen *ebiten.Image) {
	for _, entity := range sb.entities {
		entity.Draw(screen)
	}
}

func (sb *SceneBase) initEntities() {
	for _, entity := range sb.addedEntities {
		entity.Init()
	}
	sb.entities = append(sb.entities, sb.addedEntities...)
	sb.addedEntities = make([]entity.Entity, 0)
}

func (sb *SceneBase) removeDisposedEntities() {
	var aliveEntities []entity.Entity
	for _, entity := range sb.entities {
		if entity.IsActive() {
			aliveEntities = append(aliveEntities, entity)
		} else {
			entity.OnDisposed()
		}
	}
	sb.entities = aliveEntities
}
