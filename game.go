package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sijiaoh/ebiten_shooting/camera"
	"github.com/sijiaoh/ebiten_shooting/game"
	"github.com/sijiaoh/ebiten_shooting/scenes/world"
	"github.com/sijiaoh/ebiten_shooting/time"
)

type Game struct{}

func (g *Game) Init() {
	ebiten.SetWindowSize(camera.ScreenWidth, camera.ScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")

	game.C.Scene = world.NewScene()
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
