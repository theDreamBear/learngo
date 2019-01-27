package main

import "fmt"

func myDefer() {
	fmt.Println("hello world")
}

func main() {
	defer fmt.Println("end")
	defer myDefer()
	fmt.Println("start")

}
