package rules

type Simpson struct {
}

func (Simpson) IntegrateOnSegment(a, b float64, f func(float64) float64) float64 {
	coeff := (b - a) / 6.0
	xZero := (b + a) / 2.0

	return coeff * (f(a) + 4.0*f(xZero) + f(b))
}
