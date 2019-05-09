package main

import (
	"fmt"
	"unsafe"
)

// 切片前开后闭
func sliceInit() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[2:6]
	fmt.Println(s)
}

func updateSlice(s []int) {
	s[0] = 100
}

//func sliceslice() {
//	sce := []int{1, 2, 3, 4, 5, 6}
//	printSlice(sce)
//	s := sce[1:5]
//	printSlice(s)
//}

func makeslice() {
	var a = [3]byte{'a', 'b', 'c'}
	var ptr uintptr = uintptr(unsafe.Pointer(&a))
	var length = 40
	var s1 = struct {
		addr uintptr
		len  int
		cap  int
	}{ptr, length, length}

	s := *(*[]byte)(unsafe.Pointer(&s1))
	fmt.Println(len(s))
	fmt.Println(string(s[0]))
	//fmt.Println(cap(s))

}

func main() {
	//sliceInit()
	//arr := [...]int{0,1,2,3,4,5,6,7,8,9}
	//s := arr[1:7]
	//fmt.Println(s)
	//updateSlice(s)
	//fmt.Println(s)
	//sliceslice()
	makeslice()
}
