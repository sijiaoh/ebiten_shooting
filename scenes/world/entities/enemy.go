package entities

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/camera"
	"github.com/sijiaoh/ebiten_shooting/core"
)

type Enemy struct {
	EntityBase

	speedPerSecond float64
}

func NewEnemy(playerEntity *PlayerEntity) *Enemy {
	e := &Enemy{
		EntityBase:     *NewEntityBase(),
		speedPerSecond: 1,
	}
	e.AddComponent(NewFollower(e, playerEntity))
	return e
}

func (e *Enemy) Awake() {
}

func (e *Enemy) Init() {
}

func (e *Enemy) Update(delta float64) {
}

func (e *Enemy) Draw(dm *core.DrawerManager) {
	dm.AddDrawer(&core.Drawer{
		Draw: func(screen *ebiten.Image) {
			size := 0.5 * camera.PixelsPerUnit
			screenPos := camera.ToScreenPos(e.pos)
			leftTop := screenPos.Add(gmath.Vec{X: -size / 2, Y: -size / 2})
			vector.DrawFilledRect(screen, float32(leftTop.X), float32(leftTop.Y), float32(size), float32(size), color.RGBA{255, 0, 0, 255}, false)
		},
	})
}

func (e *Enemy) OnDisposed() {
}

func (e *Enemy) SpeedPerSecond() float64 {
	return e.speedPerSecond
}
