package manager

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sijiaoh/ebiten_shooting/units"
)

type ActorManager struct {
	actors []units.Actor
}

var am = &ActorManager{}

// OnDeadを確実に呼び出すためにもRemoveUnitは提供しない
func AddActor(actor units.Actor) {
	am.actors = append(am.actors, actor)
}

func Update() {
	am.initActors()

	for _, actor := range am.actors {
		if actor.IsAlive() {
			actor.Update()
		}
	}

	am.removeDeadActors()
}

func Draw(screen *ebiten.Image) {
	for _, actor := range am.actors {
		actor.Draw(screen)
	}
}

func (am *ActorManager) initActors() {
	for _, actor := range am.actors {
		if actor.IsAlive() && !actor.IsInited() {
			actor.Init()
			actor.EndInit()
		}
	}
}

func (am *ActorManager) removeDeadActors() {
	var aliveActors []units.Actor
	for _, actor := range am.actors {
		if actor.IsAlive() {
			aliveActors = append(aliveActors, actor)
		} else {
			actor.OnDead()
		}
	}
	am.actors = aliveActors
}
