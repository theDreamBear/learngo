package main

import "fmt"

func myDefer() {
	fmt.Println("hello world")
}

func catch() {
	fmt.Println("catch: ", recover())
}

func saveCode(x, y int) {
	z := 0
	defer func() {
		if recover() != nil {
			z = 0
		}
	}()
	z = x / y
	fmt.Println(z)
}

func deferInvoke() {
	defer catch()
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		recover()
	}()
	panic("i am dead")
}

func main() {
	saveCode(1, 0)
}
