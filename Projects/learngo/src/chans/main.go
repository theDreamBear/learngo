package main

import (
	"chans/pipeline"
	"fmt"
)

func main() {
	p := pipeline.Merge(pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(pipeline.ArraySource(10, 5, 8, 3, 1, 9)))
	for v := range p {
		fmt.Println(v)
	}
}
