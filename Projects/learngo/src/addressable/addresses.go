package main

import "fmt"

//func addressable_() {
//	var x = 10
//	fmt.Println(&x)
//
//	//var strmap = map[string]int{"a": 1}
//	//fmt.Println(&strmap["a"])
//
//	var sliceInt = []int{1, 2, 3}
//	fmt.Println(&sliceInt[0])
//}

func test() *int {
	a := 100
	return &a
}

func main() {
	//addressable_()
	var a *int = test()
	fmt.Println(a)
}
