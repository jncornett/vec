package vec

import "math/cmplx"

type Mass struct {
	Sum Value
	N   int
}

func Combine(m, n Mass) Mass {
	return Mass{m.Sum + n.Sum, m.N + n.N}
}

func Subtract(m, n Mass) Mass {
	return Mass{m.Sum - n.Sum, m.N - n.N}
}

func (m Mass) Center() Value {
	if m.N <= 0 {
		return Value(cmplx.NaN())
	}
	return m.Sum.Scale(1 / float64(m.N))
}
