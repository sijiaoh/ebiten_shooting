package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/camera"
	"github.com/sijiaoh/ebiten_shooting/game"
	"github.com/sijiaoh/ebiten_shooting/scenes/world/actors/enemies"
	"github.com/sijiaoh/ebiten_shooting/scenes/world/actors/player"
	"github.com/sijiaoh/ebiten_shooting/time"
)

type Game struct{}

func (g *Game) Init() {
	ebiten.SetWindowSize(camera.ScreenWidth, camera.ScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")

	playerActor := player.NewPlayerActor()
	playerActor.Pos = camera.ToWorldPos(gmath.Vec{camera.ScreenWidth / 2, camera.ScreenHeight / 2})
	game.C.Scene.AddActor(&playerActor)

	enemy := enemies.NewEnemy()
	enemy.Pos = camera.ToWorldPos(gmath.Vec{camera.ScreenWidth / 2, 1 * camera.PixelsPerUnit})
	game.C.Scene.AddActor(&enemy)
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	time.Time.OnBeforeUpdate()
	game.C.Scene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	game.C.Scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return camera.ScreenWidth, camera.ScreenHeight
}
