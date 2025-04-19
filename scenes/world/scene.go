package world

import (
	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/camera"
	"github.com/sijiaoh/ebiten_shooting/scenes"
	"github.com/sijiaoh/ebiten_shooting/scenes/world/actors/enemies"
	"github.com/sijiaoh/ebiten_shooting/scenes/world/actors/player"
)

type Scene struct {
	scenes.SceneBase
}

func NewScene() scenes.Scene {
	scene := Scene{}

	playerActor := player.NewPlayerActor()
	playerActor.Pos = camera.ToWorldPos(gmath.Vec{camera.ScreenWidth / 2, camera.ScreenHeight / 2})
	scene.AddActor(&playerActor)

	enemy := enemies.NewEnemy()
	enemy.Pos = camera.ToWorldPos(gmath.Vec{camera.ScreenWidth / 2, 1 * camera.PixelsPerUnit})
	scene.AddActor(&enemy)

	return &scene
}
