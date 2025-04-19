package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sijiaoh/ebiten_shooting/scenes/world/actors"
)

type Scene interface {
	AddActor(actor actors.Actor)
	Update()
	Draw(screen *ebiten.Image)
}

type SceneBase struct {
	actors []actors.Actor
}

// OnDisposedを確実に呼び出すためにもRemoveActorは提供しない
func (sb *SceneBase) AddActor(actor actors.Actor) {
	sb.actors = append(sb.actors, actor)
}

func (sb *SceneBase) Update() {
	sb.initActors()

	for _, actor := range sb.actors {
		if actor.IsActive() {
			actor.Update()
		}
	}

	sb.removeDeadActors()
}

func (sb *SceneBase) Draw(screen *ebiten.Image) {
	for _, actor := range sb.actors {
		actor.Draw(screen)
	}
}

func (sb *SceneBase) initActors() {
	for _, actor := range sb.actors {
		if actor.IsActive() && !actor.IsInited() {
			actor.Init()
			actor.EndInit()
		}
	}
}

func (sb *SceneBase) removeDeadActors() {
	var aliveActors []actors.Actor
	for _, actor := range sb.actors {
		if actor.IsActive() {
			aliveActors = append(aliveActors, actor)
		} else {
			actor.OnDisposed()
		}
	}
	sb.actors = aliveActors
}
