package vec

import (
	"math"
	"math/cmplx"
)

type Rect [2]Value

func (r Rect) Top() float64    { return r[0].Y() }
func (r Rect) Left() float64   { return r[0].X() }
func (r Rect) Bottom() float64 { return r[1].Y() }
func (r Rect) Right() float64  { return r[1].X() }
func (r Rect) NW() Value       { return r[0] }
func (r Rect) SW() Value       { return NewValue(r[0].X(), r[1].Y()) }
func (r Rect) SE() Value       { return r[1] }
func (r Rect) NE() Value       { return NewValue(r[1].X(), r[0].Y()) }

func (r Rect) IsValid() bool {
	return r[0].X() <= r[1].X() &&
		r[0].Y() <= r[1].Y()
}

func (r Rect) Dim() Value {
	if !r.IsValid() {
		return NaN()
	}
	return NewValue(
		math.Max(r[1].X()-r[0].X(), 0),
		math.Max(r[1].Y()-r[0].Y(), 0),
	)
}

func (r Rect) Area() float64 {
	dim := r.Dim()
	return dim.X() * dim.Y()
}

func (r Rect) Touches(v Value) bool {
	return r.IsValid() &&
		PointwiseLessEquals(r[0], v) &&
		PointwiseLessEquals(v, r[1])
}

func (r Rect) Contains(v Value) bool {
	return r.IsValid() &&
		PointwiseLess(r[0], v) &&
		PointwiseLess(v, r[1])
}

func Intersection(r, s Rect) Rect {
	return Rect{PointwiseMax(r[0], s[0]), PointwiseMax(r[1], s[1])}
}

func Intersects(r, s Rect) bool {
	return Intersection(r, s).Area() > 0
}

func Touches(r, s Rect) bool {
	return Intersection(r, s).Area() >= 0
}

func Inf() Rect {
	return Rect{Value(-cmplx.Inf()), Value(cmplx.Inf())}
}
