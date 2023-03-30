package local_integral

type Integral struct {
	F     func(float64) float64
	QRule QuadratureRule
}

type QuadratureRule interface {
	IntegrateOnSegment(float64, float64, func(float64) float64) float64
}

func (i Integral) Integrate(a, b float64, n uint64) float64 {
	h := (b - a) / float64(n)
	sum := float64(0.0)

	for k := uint64(0); k < n; k++ {
		sum += i.QRule.IntegrateOnSegment(a, a+h, i.F)
		a += h
	}

	return sum
}
