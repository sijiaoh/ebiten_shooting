package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sijiaoh/ebiten_shooting/actors"
	"github.com/sijiaoh/ebiten_shooting/actors/enemies"
	"github.com/sijiaoh/ebiten_shooting/actors/player"
	"github.com/sijiaoh/ebiten_shooting/time"
	"github.com/sijiaoh/ebiten_shooting/utils"
)

type Game struct{}

func (g *Game) Init() {
	ebiten.SetWindowSize(utils.WindowWidth, utils.WindowHeight)
	ebiten.SetWindowTitle("Hello, World!")

	playerActor := player.NewPlayerActor()
	actors.Actors.AddActor(&playerActor)

	enemy := enemies.NewEnemy()
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
	return utils.WindowWidth / 2, utils.WindowHeight / 2
}
