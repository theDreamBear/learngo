package main

import (
	"fmt"
	"time"
)

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

func TimeStampToYearMonthDay(stime int64) int32 {
	tm := time.Unix(stime, 0)
	fmt.Println(tm.Year())
	return 0
}

func main() {
	TimeStampToYearMonthDay(1557325677)
}
