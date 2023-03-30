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

func init() {
	flag.StringVar(&ruleName, "method", "Gauss", "name of Quadrature Rule")
	flag.StringVar(&functionName, "func", "Polynom", "name of function for integrating")
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

	result := integral.Integrate(0, 2, 10)

	fmt.Println(result)
}
