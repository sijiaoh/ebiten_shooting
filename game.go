package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/actors"
	"github.com/sijiaoh/ebiten_shooting/actors/enemies"
	"github.com/sijiaoh/ebiten_shooting/actors/player"
	"github.com/sijiaoh/ebiten_shooting/camera"
	"github.com/sijiaoh/ebiten_shooting/time"
)

type Game struct{}

func (g *Game) Init() {
	ebiten.SetWindowSize(camera.ScreenWidth, camera.ScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")

	playerActor := player.NewPlayerActor()
	playerActor.Pos = camera.ToWorldPos(gmath.Vec{camera.ScreenWidth / 2, camera.ScreenHeight / 2})
	actors.ActorManager.AddActor(&playerActor)

	enemy := enemies.NewEnemy()
	enemy.Pos = camera.ToWorldPos(gmath.Vec{camera.ScreenWidth / 2, 1 * camera.PixelsPerUnit})
	actors.ActorManager.AddActor(&enemy)
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	time.Time.OnBeforeUpdate()
	actors.ActorManager.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	actors.ActorManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return camera.ScreenWidth, camera.ScreenHeight
}
