package main

import (
	"fmt"
)

type Myreader string

func (Reader Myreader) Read(p []byte) (n int, err error) {
	//p = ([]byte)(string(Reader))
	//p = append(p, []byte(string(Reader))...)
	copy(p, []byte(string(Reader)))
	//fmt.Printf("%p\n", p)
	return len(string(Reader)), nil
}

func main() {
	var Mr Myreader = "hello world"
	var Mb = make([]byte, 15)
	if i, err := Mr.Read(Mb); err == nil {
		fmt.Printf("read\t%v, content is %s", i, string(Mb[:i]))
	}
	//fmt.Printf("%p", Mb)
}
