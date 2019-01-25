package main

import (
	"errors"
)

// 多返回函数
func div(a, b int)(int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a/b, nil
}

func funDefer(){
	println("close ok")
}

// 延长调用 多个defer 先声明后调用
func testDefer(a , b int) {
	defer funDefer()
	if a > 1{
		return
	} else {
		println("hello world")
	}
}

func main() {
	//a, b := 10, 2
	//c, err := div(a, b)
	//fmt.Println(c, err)

	testDefer(1,4)
}