package world

import (
	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/camera"
	"github.com/sijiaoh/ebiten_shooting/core"
	"github.com/sijiaoh/ebiten_shooting/scenes/world/actors"
)

type Scene struct {
	core.SceneBase
}

func NewScene() *Scene {
	scene := &Scene{
		SceneBase: core.NewSceneBase(),
	}

	playerActor := actors.NewPlayerActor()
	playerActor.SetPos(camera.ToWorldPos(gmath.Vec{camera.ScreenWidth / 2, camera.ScreenHeight / 2}))
	scene.AddEntity(playerActor)

	enemy := actors.NewEnemy(playerActor)
	enemy.SetPos(camera.ToWorldPos(gmath.Vec{camera.ScreenWidth / 2, 1 * camera.PixelsPerUnit}))
	scene.AddEntity(enemy)

	return scene
}
