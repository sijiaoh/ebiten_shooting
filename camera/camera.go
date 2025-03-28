package camera

import "github.com/sijiaoh/ebiten_shooting/utils"

const (
	PixelsPerUnit = 32
	ScreenWidth   = 640
	ScreenHeight  = 480
)

func ToScreenPos(pos utils.VectorFloat) utils.VectorFloat {
	return utils.VectorFloat{pos.X * PixelsPerUnit, pos.Y * PixelsPerUnit}
}

func ToWorldPos(pos utils.VectorFloat) utils.VectorFloat {
	return utils.VectorFloat{pos.X / PixelsPerUnit, pos.Y / PixelsPerUnit}
}
