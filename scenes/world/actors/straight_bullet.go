package actors

import (
	"image/color"

	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/core"
)

type StraightBullet struct {
	ActorBase

	direction gmath.Vec
	speed     float64
}

func NewStraightBullet(pos gmath.Vec, direction gmath.Vec, speed float64) *StraightBullet {
	sb := &StraightBullet{
		ActorBase: *NewActorBase(),
		direction: direction,
		speed:     speed,
	}
	sb.pos = pos
	return sb
}

func (sb *StraightBullet) Awake() {
}

func (sb *StraightBullet) Init() {
}

func (sb *StraightBullet) Update(delta float64) {
	speed := sb.speed * delta
	sb.pos = sb.pos.Add(sb.direction.Mulf(speed))
}

func (sb *StraightBullet) Draw(dm *core.DrawerManager) {
	drawer := dm.NewDrawer()
	drawer.DrawCircle(sb.pos, 0.1, color.RGBA{R: 255, G: 255, B: 0, A: 255})
	dm.AddDrawer(drawer)
}

func (sb *StraightBullet) OnDisposed() {
}
