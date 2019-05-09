package main

import "fmt"

func rawString() {
	var s = `ssdd
	sdfsf
fgdsfs
dsfsgfs
`
	fmt.Println(s)
}

func stringAdd() {
	s := "ab" +
		"cd"

	fmt.Println(s == "abcd")
	fmt.Println(s > "abcd")
}

func stringModify() {
	//var s = "hello"
	//var dd  = s[:]
	//dd[0] = 'f'

	var s = []byte{1, 2, 3}
	fmt.Println(&s, s)
	s[0] = 4
	fmt.Println(&s, s)
}

func main() {
	rawString()
	stringAdd()
	stringModify()
}
