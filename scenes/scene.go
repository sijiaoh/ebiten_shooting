package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sijiaoh/ebiten_shooting/scenes/world/entities"
)

type Scene interface {
	AddEntity(entity entities.Entity)
	Update()
	Draw(screen *ebiten.Image)
}

type SceneBase struct {
	entities []entities.Entity
}

// OnDisposedを確実に呼び出すためにもRemoveEntityは提供しない
func (sb *SceneBase) AddEntity(entity entities.Entity) {
	sb.entities = append(sb.entities, entity)
}

func (sb *SceneBase) Update() {
	sb.initEntities()

	for _, entity := range sb.entities {
		if entity.IsActive() {
			entity.Update()
		}
	}

	sb.removeDeadEntities()
}

func (sb *SceneBase) Draw(screen *ebiten.Image) {
	for _, entity := range sb.entities {
		entity.Draw(screen)
	}
}

func (sb *SceneBase) initEntities() {
	for _, entity := range sb.entities {
		if entity.IsActive() && !entity.IsInited() {
			entity.Init()
			entity.EndInit()
		}
	}
}

func (sb *SceneBase) removeDeadEntities() {
	var aliveEntities []entities.Entity
	for _, entity := range sb.entities {
		if entity.IsActive() {
			aliveEntities = append(aliveEntities, entity)
		} else {
			entity.OnDisposed()
		}
	}
	sb.entities = aliveEntities
}
