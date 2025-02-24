package time

type TimeManager struct {
	DeltaTime          float64
	prevFrameStartTime float64
}

var Time = TimeManager{}

func (t *TimeManager) OnBeforeUpdate() {
	// Ebitenは60フレームを前提にするとのこと
	t.DeltaTime = 1.0 / 60.0
}
