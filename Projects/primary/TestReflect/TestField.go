package main

import (
	"fmt"
	"reflect"
)

func reflect_test() {
	var s  = make(chan string)
	obj := reflect.ValueOf(s)

	fmt.Println(obj.Field(1))
}


func main() {

}