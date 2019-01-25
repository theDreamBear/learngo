package main

import "fmt"

func test1(s string, a... int){
	fmt.Printf("%T, %v",a,a)
}


func main() {
	test1("abc", 1,2,3,4)
}
