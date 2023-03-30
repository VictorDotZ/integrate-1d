package functions

import "math"

func Polynom(x float64) float64 {
	return 6.0 * math.Pow(x, 5)
}

func Cos100(x float64) float64 {
	return math.Cos(100.0 * x)
}

func ExpMinus1000(x float64) float64 {
	return math.Exp(x * (-1000.0))
}

func Density(x float64) float64 {
	return 1.0 / math.Sqrt(1.0-x*x)
}
