package main

import (
	"fmt"
	"functional/fabnocci"
)

func main() {
	var fn = fabnocci.Fabnocci()
	for i := 0; i < 10; i++ {
		fmt.Println(fn())
	}
}
