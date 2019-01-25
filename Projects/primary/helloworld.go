package main

import (
	"fmt"
	"time"
)

func main(){
	for i := 0; i < 50000 ; i++ {
		go helloWorld(i)
	}
	time.Sleep(time.Microsecond * 30)
}


func helloWorld(i int) {
	fmt.Printf("%d: helloworld\n", i)
}