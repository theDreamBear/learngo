package main

import "fmt"

func initArray() {
	var d1 [3]int
	fmt.Println(d1)

	var d2 = [4]int{1,2}
	fmt.Println(d2)

	var d3 = [...]int{1,2,3}
	fmt.Println(d3)

	var d4 = [...]int{10,3:20}
	fmt.Println(d4)
}

func sliceType() {
	s := "test"
	xByte := []byte(s)
	fmt.Printf("%T\n",xByte)
	xByte[0] = 'T'
	s1 := string(xByte)
	fmt.Println(s1)
}

func main () {
	//initArray()
	sliceType()
}
