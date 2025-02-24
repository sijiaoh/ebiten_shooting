package actors

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ActorManager struct {
	actors []Actor
}

var Actors = &ActorManager{}

// OnDeadを確実に呼び出すためにもRemoveActorは提供しない
func (am *ActorManager) AddActor(actor Actor) {
	am.actors = append(am.actors, actor)
}

func (am *ActorManager) Update() {
	am.initActors()

	for _, actor := range am.actors {
		if actor.IsAlive() {
			actor.Update()
		}
	}

	am.removeDeadActors()
}

func (am *ActorManager) Draw(screen *ebiten.Image) {
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
	var aliveActors []Actor
	for _, actor := range am.actors {
		if actor.IsAlive() {
			aliveActors = append(aliveActors, actor)
		} else {
			actor.OnDead()
		}
	}
	am.actors = aliveActors
}
