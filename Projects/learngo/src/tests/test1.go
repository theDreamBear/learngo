package main

import "fmt"

var p *int

func foo() (*int, error) {
	var i int = 5
	fmt.Printf("%p\n", &i)
	return &i, nil
}

func bar() {
	//use p
	//fmt.Printf("%p\n", p)
	fmt.Println(*p)
}

func main() {
	//fmt.Printf("%p\n", p)
	var err error
	p, err = foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}
