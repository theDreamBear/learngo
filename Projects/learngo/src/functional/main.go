package main

import (
	"bufio"
	"fmt"
	"io"
)

func printFileContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func innerFunc(i int) {
	fmt.Println("hello world")
}

func funcWrapper(f func(i int), i int) {
	f(i)
}

type Fabs func() int

func (f *Fabs) Next() int {
	return (*f)()
}

func Fabnocci() Fabs {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main1() {
	//var fn = fabnocci.Fabnocci()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(fn())
	//}
	//
	//var fn1 = fabnocci.Fabnocci()
	//printFileContent(fn1)
	//
	//

}

func main() {
	funcWrapper(innerFunc, 10)
}
