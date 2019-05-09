package main

import (
	"bufio"
	"fmt"
	"strings"
)

func stringScanner() {
	var names = "nic silent tom"
	var scannner = bufio.NewScanner(strings.NewReader(names))
	scannner.Split(bufio.ScanWords)
	for scannner.Scan() {
		fmt.Println(scannner.Text())
	}
}

func main() {
	stringScanner()
}
