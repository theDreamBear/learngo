package main

import (
	"duckTyping/calculator"
	"fmt"
)

func main() {
	var ca calculator.Calculator
	ca = calculator.NewCalculator{}
	fmt.Println(ca.Div(float64(3), float64(12)))

	ca = calculator.ComplexCalculator{}
	fmt.Println(ca.Div(calculator.Mcomplex{1, 2}, calculator.Mcomplex{2, 1}))
}
