package main

import (
	"fmt"
	"strconv"
)

func varFormat(){
	a,b := 100 ,01234
	fmt.Printf("a is %d\n", a)
	fmt.Printf("b is ob%b, %#o, %#x\n", b,b,b)
}

func 进制转换() {
	// strconv  --> stringConversion
	a, ok := strconv.ParseInt("1101",3,32)
	if ok == nil {
		println("a is", a)
	} else {
		panic(ok)
	}
}

func main() {
	//varFormat()
	进制转换()
}

