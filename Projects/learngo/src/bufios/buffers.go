package main

import "bytes"

type buffer struct {
	bytes.Buffer
	tmp  [64]byte // temporary byte array for creating headers.
	next *buffer
}

func main() {

}
