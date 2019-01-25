package main

import "fmt"

func Test(x int) func(){
	println(&x)
	return func() {
		println(&x, x)
	}
}

func Test1() []func(){
	var s[] func()
	for i := 0; i < 3; i++ {
		s = append(s, func(){
			println(&i, i)
		})
	}
	return s
}

func main() {
	x := 123
	f := Test(x)
	x += 1
	f()
	fmt.Println("++++++++++++++++++++++++++++++++++")

	for _, f := range Test1() {
		f()
	}
}
