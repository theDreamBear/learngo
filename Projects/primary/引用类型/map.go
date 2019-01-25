package main

import "fmt"

func TestMap() map[int]string{
	mp := make(map[int]string)
	mp[0] = "nihao"
	mp[10] = "zai"
	return mp
}

func main() {
	mp := TestMap()
	for idx, name := range(mp){
		println(idx, name)
	}

	var f = func(s string) bool{
		println(s)
		println("hello world")
		return false
	}
	f("good")

	type user struct{
		x int
		name string
	}

	n := user{10000000, "nicj",}

	fmt.Println(n.x,n.name)
}