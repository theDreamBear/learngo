package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBinary(number int) string {
	result := ""
	for; number != 0; {
		md := number % 2;
		number /= 2
		result = strconv.Itoa(md) + result
	}
	return result
}

func printFile(path string){
	if file, e := os.Open(path); e != nil {
		panic(e)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}

func main() {
	const file = "./abc.txt"
	fmt.Println(convertToBinary(13))
	printFile(file)
}
