package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sijiaoh/ebiten_shooting/scenes/world/actors"
)

type Scene struct {
	actors []actors.Actor
}

// OnDeadを確実に呼び出すためにもRemoveActorは提供しない
func (am *Scene) AddActor(actor actors.Actor) {
	am.actors = append(am.actors, actor)
}

func (am *Scene) Update() {
	am.initActors()

	for _, actor := range am.actors {
		if actor.IsAlive() {
			actor.Update()
		}
	}

	am.removeDeadActors()
}

func (am *Scene) Draw(screen *ebiten.Image) {
	for _, actor := range am.actors {
		actor.Draw(screen)
	}
}

func (am *Scene) initActors() {
	for _, actor := range am.actors {
		if actor.IsAlive() && !actor.IsInited() {
			actor.Init()
			actor.EndInit()
		}
	}
}

func (am *Scene) removeDeadActors() {
	var aliveActors []actors.Actor
	for _, actor := range am.actors {
		if actor.IsAlive() {
			aliveActors = append(aliveActors, actor)
		} else {
			actor.OnDead()
		}
	}
	am.actors = aliveActors
}
