package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sijiaoh/ebiten_shooting/actors"
	"github.com/sijiaoh/ebiten_shooting/time"
)

type Game struct{}

func (g *Game) Init() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	playerActor := actors.NewPlayerActor()
	actors.Actors.AddActor(&playerActor)
}

func (g *Game) Update() error {
	time.Time.OnBeforeUpdate()
	actors.Actors.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	actors.Actors.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
