package utils

import (
	"math"
)

type VectorFloat struct {
	X float64
	Y float64
}

func VectorFloatZero() VectorFloat {
	return VectorFloat{0, 0}
}

func VectorFloatUp() VectorFloat {
	return VectorFloat{0, -1}
}

func VectorFloatDown() VectorFloat {
	return VectorFloat{0, 1}
}

func VectorFloatLeft() VectorFloat {
	return VectorFloat{-1, 0}
}

func VectorFloatRight() VectorFloat {
	return VectorFloat{1, 0}
}

func (vf *VectorFloat) Add(v VectorFloat) VectorFloat {
	return VectorFloat{vf.X + v.X, vf.Y + v.Y}
}

func (vf *VectorFloat) Sub(v VectorFloat) VectorFloat {
	return VectorFloat{vf.X - v.X, vf.Y - v.Y}
}

func (vf *VectorFloat) Mul(s float64) VectorFloat {
	return VectorFloat{vf.X * s, vf.Y * s}
}

func (vf *VectorFloat) Div(s float64) VectorFloat {
	return VectorFloat{vf.X / s, vf.Y / s}
}

func (vf *VectorFloat) LengthSquared() float64 {
	return vf.X*vf.X + vf.Y*vf.Y
}

func (vf *VectorFloat) Length() float64 {
	return math.Sqrt(vf.LengthSquared())
}

func (vf *VectorFloat) Normalize() VectorFloat {
	return vf.Div(vf.Length())
}
