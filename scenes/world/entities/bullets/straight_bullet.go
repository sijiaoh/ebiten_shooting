package bullets

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/quasilyte/gmath"
	"github.com/sijiaoh/ebiten_shooting/camera"
	"github.com/sijiaoh/ebiten_shooting/core"
)

type StraightBullet struct {
	core.EntityBase

	Pos       gmath.Vec
	direction gmath.Vec
	speed     float64
}

func NewStraightBullet(pos gmath.Vec, direction gmath.Vec, speed float64) *StraightBullet {
	sb := &StraightBullet{
		EntityBase: *core.NewEntityBase(),
		Pos:        pos,
		direction:  direction,
		speed:      speed,
	}
	return sb
}

func (sb *StraightBullet) Awake() {
}

func (sb *StraightBullet) Init() {
}

func (sb *StraightBullet) Update(delta float64) {
	speed := sb.speed * delta
	sb.Pos = sb.Pos.Add(sb.direction.Mulf(speed))
}

func (sb *StraightBullet) Draw(dm *core.DrawerManager) {
	dm.AddDrawer(&core.Drawer{
		Draw: func(screen *ebiten.Image) {
			size := 0.1 * camera.PixelsPerUnit
			screenPos := camera.ToScreenPos(sb.Pos)
			vector.DrawFilledCircle(screen, float32(screenPos.X), float32(screenPos.Y), float32(size/2), color.RGBA{R: 255, G: 255, B: 0, A: 255}, false)
		},
	})
}

func (sb *StraightBullet) OnDisposed() {
}
