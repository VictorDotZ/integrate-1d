package main

import (
	"flag"
	"fmt"
	"integrate1d/pkg/functions"
	"integrate1d/pkg/local_integral"
	"integrate1d/pkg/rules"

	"os"
)

func usage() {
	fmt.Println("usage:")
	flag.PrintDefaults()
	os.Exit(2)
}

var ruleName string
var functionName string
var a float64
var b float64
var n uint64

func init() {
	flag.StringVar(&ruleName, "method", "Gauss", "name of Quadrature Rule")
	flag.StringVar(&functionName, "func", "Polynom", "name of function for integrating")
	flag.Float64Var(&a, "a", 0.0, "start of segment")
	flag.Float64Var(&b, "b", 1.0, "end of segment")
	flag.Uint64Var(&n, "n", 100, "number of splits")

}

func main() {
	flag.Usage = usage
	flag.Parse()

	functionNames := map[string]func(float64) float64{
		"Polynom":      functions.Polynom,
		"Cos100":       functions.Cos100,
		"ExpMinus1000": functions.ExpMinus1000,
		"Density":      functions.Density,
	}

	rules := map[string]local_integral.QuadratureRule{
		"Gauss":   rules.Gauss{},
		"Simpson": rules.Simpson{},
	}

	if _, ok := functionNames[functionName]; !ok {
		if _, ok := rules[ruleName]; !ok {
			usage()
		}
	}

	integral := local_integral.Integral{
		F:     functionNames[functionName],
		QRule: rules[ruleName],
	}

	result := integral.Integrate(a, b, n)

	fmt.Println(result)
}
