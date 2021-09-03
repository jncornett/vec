package vec

import (
	"math"
	"math/cmplx"
)

type Value complex128

func (v Value) X() float64 { return real(v) }
func (v Value) Y() float64 { return imag(v) }

func (v Value) R() float64 {
	r, _ := cmplx.Polar(complex128(v))
	return r
}

func (v Value) Θ() float64 {
	_, θ := cmplx.Polar(complex128(v))
	return θ
}

func (v Value) Scale(scalar float64) Value {
	return Value(complex(real(v)*scalar, imag(v)*scalar))
}

func (v Value) Unit() Value { return v.Scale(1 / v.R()) }

func (v Value) IsNaN() bool {
	return cmplx.IsNaN(complex128(v))
}

func Dist(v, w Value) float64 {
	return (v - w).R()
}

func PointwiseMax(v, w Value) Value {
	return Value(complex(math.Max(real(v), real(w)), math.Max(imag(v), imag(w))))
}

func PointwiseMin(v, w Value) Value {
	return Value(complex(math.Min(real(v), real(w)), math.Min(imag(v), imag(w))))
}

func PointwiseRemainder(v, w Value) Value {
	return Value(complex(math.Remainder(real(v), real(w)), math.Remainder(imag(v), imag(w))))
}

func PointwiseLess(v, w Value) bool {
	return real(v) < real(w) && imag(v) < imag(w)
}

func PointwiseLessEquals(v, w Value) bool {
	return real(v) <= real(w) && imag(v) <= imag(w)
}

func PointwiseGreater(v, w Value) bool {
	return real(v) > real(w) && imag(v) > imag(w)
}

func PointwiseGreaterEquals(v, w Value) bool {
	return real(v) >= real(w) && imag(v) >= imag(w)
}

func NewValue(re, im float64) Value {
	return Value(complex(re, im))
}

func Real(scalar float64) Value {
	return Value(complex(scalar, 0))
}

func Imag(scalar float64) Value {
	return Value(complex(0, scalar))
}

func Scalar(scalar float64) Value {
	return Value(complex(scalar, scalar))
}

func NaN() Value {
	return Value(cmplx.NaN())
}
