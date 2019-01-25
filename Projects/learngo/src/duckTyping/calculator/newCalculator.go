package calculator

import "fmt"

type NewCalculator struct {
}

func (nc NewCalculator) Add(a, b interface{}) interface{} {
	//fmt.Printf("%T\t%v\n", a, a)
	_print()
	return a.(float64) + b.(float64)
}

func (nc NewCalculator) Sub(a, b interface{}) interface{} {
	_print()
	return a.(float64) - b.(float64)
}

func (nc NewCalculator) Mul(a, b interface{}) interface{} {
	_print()
	return a.(float64) * b.(float64)
}

func (nc NewCalculator) Div(a, b interface{}) interface{} {
	_print()
	if b != float64(0) {
		return a.(float64) / b.(float64)
	} else {
		panic("error: divide by zero")
	}
}

func _print() {
	fmt.Printf("NOTICE: calculate by NewCalculator, the answer is: ")
}
