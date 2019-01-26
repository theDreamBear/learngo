package main

import (
	"bufio"
	"fmt"
	"functional/fabnocci"
	"io"
)

func printFileContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	//var fn = fabnocci.Fabnocci()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(fn())
	//}

	var fn1 = fabnocci.Fabnocci()
	printFileContent(fn1)

}
