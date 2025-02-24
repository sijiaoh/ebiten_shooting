package utils

import (
	"math"
)

type VectorFloat struct {
	x float64
	y float64
}

func VectorFloatZero() VectorFloat {
	return VectorFloat{0, 0}
}

func (vf *VectorFloat) Add(v VectorFloat) VectorFloat {
	return VectorFloat{vf.x + v.x, vf.y + v.y}
}

func (vf *VectorFloat) Sub(v VectorFloat) VectorFloat {
	return VectorFloat{vf.x - v.x, vf.y - v.y}
}

func (vf *VectorFloat) Mul(s float64) VectorFloat {
	return VectorFloat{vf.x * s, vf.y * s}
}

func (vf *VectorFloat) Div(s float64) VectorFloat {
	return VectorFloat{vf.x / s, vf.y / s}
}

func (vf *VectorFloat) LengthSquared() float64 {
	return vf.x*vf.x + vf.y*vf.y
}

func (vf *VectorFloat) Length() float64 {
	return math.Sqrt(vf.LengthSquared())
}

func (vf *VectorFloat) Normalize() VectorFloat {
	return vf.Div(vf.Length())
}
