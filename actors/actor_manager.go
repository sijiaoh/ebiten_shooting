package actors

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Actors struct {
	actors []Actor
}

var ActorManager = &Actors{}

// OnDeadを確実に呼び出すためにもRemoveActorは提供しない
func (am *Actors) AddActor(actor Actor) {
	am.actors = append(am.actors, actor)
}

func (am *Actors) Update() {
	am.initActors()

	for _, actor := range am.actors {
		if actor.IsAlive() {
			actor.Update()
		}
	}

	am.removeDeadActors()
}

func (am *Actors) Draw(screen *ebiten.Image) {
	for _, actor := range am.actors {
		actor.Draw(screen)
	}
}

func (am *Actors) initActors() {
	for _, actor := range am.actors {
		if actor.IsAlive() && !actor.IsInited() {
			actor.Init()
			actor.EndInit()
		}
	}
}

func (am *Actors) removeDeadActors() {
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
