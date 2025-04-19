package enemies

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/camera"
	"github.com/sijiaoh/ebiten_shooting/scenes/world/actors"
)

type Enemy struct {
	actors.ActorBase

	Pos            gmath.Vec
	speedPerSecond float64
}

func NewEnemy() Enemy {
	return Enemy{
		ActorBase:      actors.NewActorBase(),
		speedPerSecond: 1,
	}
}

func (e *Enemy) Init() {
}

func (e *Enemy) Update() {
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	size := 0.5 * camera.PixelsPerUnit
	screenPos := camera.ToScreenPos(e.Pos)
	leftTop := screenPos.Add(gmath.Vec{X: -size / 2, Y: -size / 2})
	vector.DrawFilledRect(screen, float32(leftTop.X), float32(leftTop.Y), float32(size), float32(size), color.White, false)
}

func (e *Enemy) OnDisposed() {
}
