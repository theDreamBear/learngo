package main

// local var escape
func test() *int{
	a := 100
	return &a
}

func main() {
	var a *int = test()
	println(a, *a)
}
