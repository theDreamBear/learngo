package calculator

import "fmt"

type Mcomplex struct {
	Real  float64
	Image float64
}

type ComplexCalculator struct{}

func (c Mcomplex) String() string {
	return fmt.Sprintf("(%f, %fi)\n", c.Real, c.Image)
}

func (cc ComplexCalculator) Add(a interface{}, b interface{}) interface{} {
	_print1()
	return Mcomplex{a.(Mcomplex).Real + b.(Mcomplex).Real, a.(Mcomplex).Image + b.(Mcomplex).Image}
}

func (cc ComplexCalculator) Sub(a interface{}, b interface{}) interface{} {
	_print1()
	return Mcomplex{a.(Mcomplex).Real - b.(Mcomplex).Real, a.(Mcomplex).Image - b.(Mcomplex).Image}
}

func (cc ComplexCalculator) Mul(a interface{}, b interface{}) interface{} {
	_print1()
	return Mcomplex{a.(Mcomplex).Real*b.(Mcomplex).Real - a.(Mcomplex).Image*b.(Mcomplex).Image,
		a.(Mcomplex).Image*b.(Mcomplex).Real + a.(Mcomplex).Real*b.(Mcomplex).Image}
}

func (cc ComplexCalculator) Div(a interface{}, b interface{}) interface{} {
	_print1()
	if b.(Mcomplex).Real == float64(0) && b.(Mcomplex).Image == float64(0) {
		panic("error: divide by zero")
	} else {
		var denominatator = (b.(Mcomplex).Real*b.(Mcomplex).Real + b.(Mcomplex).Image*b.(Mcomplex).Image)
		return Mcomplex{(a.(Mcomplex).Real*b.(Mcomplex).Real + a.(Mcomplex).Image*b.(Mcomplex).Image) / denominatator,
			(a.(Mcomplex).Image*b.(Mcomplex).Real - a.(Mcomplex).Real*b.(Mcomplex).Image) / denominatator}
	}
}

func _print1() {
	fmt.Printf("NOTICE: calculate by ComplexCalculator, the answer is: ")

}
