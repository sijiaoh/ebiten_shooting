package camera

import "github.com/quasilyte/gmath"

const (
	PixelsPerUnit = 32
	ScreenWidth   = 640
	ScreenHeight  = 480
)

func ToScreenPos(pos gmath.Vec) gmath.Vec {
	return gmath.Vec{pos.X * PixelsPerUnit, pos.Y * PixelsPerUnit}
}

func ToWorldPos(pos gmath.Vec) gmath.Vec {
	return gmath.Vec{pos.X / PixelsPerUnit, pos.Y / PixelsPerUnit}
}
