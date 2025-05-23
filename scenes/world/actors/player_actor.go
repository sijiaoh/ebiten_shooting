package actors

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/camera"
	"github.com/sijiaoh/ebiten_shooting/core"
	"github.com/sijiaoh/ebiten_shooting/game"
)

type PlayerActor struct {
	ActorBase

	speedPerSecond float64
}

func NewPlayerActor() *PlayerActor {
	return &PlayerActor{
		ActorBase:      *NewActorBase(),
		speedPerSecond: 2,
	}
}

func (pa *PlayerActor) Awake() {
}

func (pa *PlayerActor) Init() {
}

func (pa *PlayerActor) Update(delta float64) {
	pa.move(delta)
	pa.shoot()
}

func (pa *PlayerActor) Draw(dm *core.DrawerManager) {
	drawer := dm.NewDrawer()
	drawer.Draw = func(screen *ebiten.Image) {
		size := 0.5 * camera.PixelsPerUnit
		screenPos := camera.ToScreenPos(pa.pos)
		vector.DrawFilledCircle(screen, float32(screenPos.X), float32(screenPos.Y), float32(size/2.0), color.White, false)
	}
	dm.AddDrawer(drawer)
}

func (pa *PlayerActor) OnDisposed() {
}

func (pa *PlayerActor) move(delta float64) {
	vec := gmath.Vec{}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		vec.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		vec.Y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		vec.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		vec.X += 1
	}

	if vec.LenSquared() > 0 {
		vec = vec.Normalized()
		speed := pa.speedPerSecond * delta
		vec = vec.Mulf(speed)
	}

	pa.pos = pa.pos.Add(vec)
}

func (pa *PlayerActor) shoot() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		b := NewStraightBullet(pa.pos, gmath.Vec{X: 0, Y: -1}, 5)
		game.C.Scene.AddEntity(b)
	}
}
