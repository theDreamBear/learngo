package main

import "fmt"

func initMap() {
	var m map[string]int
	if s, err := m["bai"]; err == false {
		fmt.Println(1)
	} else {
		fmt.Println(2, s)
	}
}

func main() {
	initMap()
}
