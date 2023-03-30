package rules

import "math"

type Gauss struct {
}

func (Gauss) IntegrateOnSegment(a, b float64, f func(float64) float64) float64 {
	coeff := (b - a) / 18.0
	xZero := (b + a) / 2.0

	shift := (b - a) / 2.0 * math.Sqrt(3.0/5.0)

	xMinus := xZero - shift
	xPlus := xZero + shift

	return coeff * (5.0*f(xMinus) + 8.0*f(xZero) + 5.0*f(xPlus))
}
