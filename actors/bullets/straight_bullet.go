package bullets

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/actors"
	"github.com/sijiaoh/ebiten_shooting/camera"
	"github.com/sijiaoh/ebiten_shooting/time"
)

type StraightBullet struct {
	actors.ActorBase

	Pos       gmath.Vec
	direction gmath.Vec
	speed     float64
}

func NewStraightBullet(pos gmath.Vec, direction gmath.Vec, speed float64) StraightBullet {
	sb := StraightBullet{
		ActorBase: actors.NewActorBase(),
		Pos:       pos,
		direction: direction,
		speed:     speed,
	}
	return sb
}

func (sb *StraightBullet) Init() {
}

func (sb *StraightBullet) Update() {
	speed := sb.speed * time.Time.DeltaTime
	sb.Pos = sb.Pos.Add(sb.direction.Mul(gmath.Vec{speed, speed}))
}

func (sb *StraightBullet) Draw(screen *ebiten.Image) {
	size := 0.1 * camera.PixelsPerUnit
	screenPos := camera.ToScreenPos(sb.Pos)
	vector.DrawFilledCircle(screen, float32(screenPos.X), float32(screenPos.Y), float32(size/2), color.RGBA{R: 255, G: 255, B: 0, A: 255}, false)
}

func (sb *StraightBullet) OnDead() {
}
