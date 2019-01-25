package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3
var bb = "string"
// cc := 1 // wrong

var (
	dd = 1
	gg = 2
)

func variableZeroValue()  {
	var(
		a int
		b string
	)
	fmt.Printf("%d %q",a, b)
	//fmt.Println(b)
}

func variableInitialValue(){
	var a, b int = 3, 4
	fmt.Println(a, b)
}

func variableTypeDeduction() {
	var a, b, c = 2, "hell", true
	fmt.Println(a, b, c)
}

func variableShort() {
	a, b, c := 2, "hell", true
	fmt.Println(a, b, c)
}

func euler() {
	fmt.Println(cmplx.Pow(math.E,1i*math.Pi) + 1)
}


func triangle() {
	var(
		a , b = 3, 4
		c int
	)
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

// 常量可提供类型，也可不提供类型
func consts() {
	const a, b = 3, 4
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	c := math.Sqrt(a * a + b * b)
	fmt.Println(c)
}

func enums() {
	const(
		c = 0
		java = 1
		python = 2
	)
	fmt.Println(c, java, python)

	const(
		mon = iota
		tue
		wed
		the
		fri
	)

	fmt.Println(mon, tue, wed,the, fri)

	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)
	fmt.Println(b, kb, mb, gb, tb)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShort()
	euler()
	triangle()
	consts()
	enums()
}
