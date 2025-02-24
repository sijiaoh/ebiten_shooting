package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sijiaoh/ebiten_shooting/units"
	unitManager "github.com/sijiaoh/ebiten_shooting/units/manager"
)

type Game struct{}

func (g *Game) Init() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	unitManager.AddUnit(&units.PlayerUnit{})
}

func (g *Game) Update() error {
	unitManager.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	unitManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
