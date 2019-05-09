package main

import "fmt"

func testChannel() {
	c := make(chan func(int, int) int, 2)
	c <- func(x, y int) int {
		return x + y
	}

	println((<-c)(1, 3))
}

func test() []func() {
	var s []func()
	for i := 0; i < 2; i++ {
		j := i
		s = append(s, func() {
			println(&j, j)
		})
	}
	return s
}

func test2(x int) func() {
	fmt.Println(&x, x)
	return func() {
		fmt.Println(&x, x)
	}
}

func main() {
	testChannel()
	for _, f := range test() {
		f()
	}
}
