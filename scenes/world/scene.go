package world

import (
	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/camera"
	"github.com/sijiaoh/ebiten_shooting/core"
	"github.com/sijiaoh/ebiten_shooting/scenes/world/entities/enemies"
	"github.com/sijiaoh/ebiten_shooting/scenes/world/entities/player"
)

type Scene struct {
	core.SceneBase
}

func NewScene() core.Scene {
	scene := Scene{
		SceneBase: core.NewSceneBase(),
	}

	playerEntity := player.NewPlayerEntity()
	playerEntity.Pos = camera.ToWorldPos(gmath.Vec{camera.ScreenWidth / 2, camera.ScreenHeight / 2})
	scene.AddEntity(playerEntity)

	enemy := enemies.NewEnemy()
	enemy.Pos = camera.ToWorldPos(gmath.Vec{camera.ScreenWidth / 2, 1 * camera.PixelsPerUnit})
	scene.AddEntity(enemy)

	return &scene
}
