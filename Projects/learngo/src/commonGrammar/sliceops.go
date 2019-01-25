package main

import (
	"errors"
	"fmt"
)

func initSlice(size int) []int {
	var sl []int
	for i :=0; i < size; i++ {
		sl = append(sl, i)
	}
	return sl
}



func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d\t",len(s), cap(s))
	fmt.Println(s)
}

// slcie 操作
func pop(s []int )([]int, error) {
	if cap(s) >= 1 {
		return s[:len(s) - 1], nil
	} else {
		return s, errors.New("切片为空")
	}
}

func front(s [] int)([]int, error) {
	if cap(s) >= 1 {
		return s[1:], nil
	} else {
		return s, errors.New("切片为空")
	}
}

func deleteSlice(s []int, i int)([]int, error) {
	if i > 0 && i < cap(s) - 1 {
		return append(s[:i],s[i + 1:]...),nil
	} else if i == 0 {
		return front(s)
	} else if i == cap(s) - 1 {
		return pop(s)
	}
	return s, errors.New("index is wrong")
}



func main() {
	s := initSlice(10)
	printSlice(s)
	if ints, e := deleteSlice(s, 9); e != nil {
		fmt.Println(e)
	} else {
		printSlice(ints)
	}
}
