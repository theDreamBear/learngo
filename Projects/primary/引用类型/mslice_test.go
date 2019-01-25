package main

import (
	"fmt"
	"reflect"
)

// 数组是值类型
func Testslice() {
	s := "hello"
	fmt.Println(reflect.TypeOf(s) ,&s, s)
	t := s[2:]
	fmt.Println(reflect.TypeOf(t) ,&t, t)

	//
	d := "hello"
	fmt.Println(reflect.TypeOf(d) ,&d, d)
	d = d[1:]
	fmt.Println(reflect.TypeOf(d) ,&d, d)


	//f := []int{1, 2, 3}

}

func main() {

}
