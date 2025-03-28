package bullets

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/sijiaoh/ebiten_shooting/actors"
	"github.com/sijiaoh/ebiten_shooting/camera"
	"github.com/sijiaoh/ebiten_shooting/time"
	"github.com/sijiaoh/ebiten_shooting/utils"
)

type StraightBullet struct {
	actors.ActorBase

	Pos       utils.VectorFloat
	direction utils.VectorFloat
	speed     float64
}

func NewStraightBullet(pos utils.VectorFloat, direction utils.VectorFloat, speed float64) StraightBullet {
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
	sb.Pos = sb.Pos.Add(sb.direction.Mul(speed))
}

func (sb *StraightBullet) Draw(screen *ebiten.Image) {
	size := 2.0
	screenPos := camera.ToScreenPos(sb.Pos)
	vector.DrawFilledCircle(screen, float32(screenPos.X), float32(screenPos.Y), float32(size/2), color.RGBA{R: 255, G: 255, B: 0, A: 255}, false)
}

func (sb *StraightBullet) OnDead() {
}
