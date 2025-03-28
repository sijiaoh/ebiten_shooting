package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sijiaoh/ebiten_shooting/actors"
	"github.com/sijiaoh/ebiten_shooting/actors/enemies"
	"github.com/sijiaoh/ebiten_shooting/actors/player"
	"github.com/sijiaoh/ebiten_shooting/camera"
	"github.com/sijiaoh/ebiten_shooting/time"
	"github.com/sijiaoh/ebiten_shooting/utils"
)

type Game struct{}

func (g *Game) Init() {
	ebiten.SetWindowSize(camera.ScreenWidth, camera.ScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")

	playerActor := player.NewPlayerActor()
	playerActor.Pos = camera.ToWorldPos(utils.VectorFloat{camera.ScreenWidth / 2, camera.ScreenHeight / 2})
	actors.Actors.AddActor(&playerActor)

	enemy := enemies.NewEnemy()
	enemy.Pos = camera.ToWorldPos(utils.VectorFloat{camera.ScreenWidth / 2, 1 * camera.PixelsPerUnit})
	actors.Actors.AddActor(&enemy)
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	time.Time.OnBeforeUpdate()
	actors.Actors.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	actors.Actors.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return camera.ScreenWidth, camera.ScreenHeight
}
