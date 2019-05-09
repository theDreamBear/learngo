package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type MyStruct struct {
	name string
	//age  int
}

var sizeOfMyStruct = int(unsafe.Sizeof(MyStruct{}))

func MyStructToBytes(s *MyStruct) []byte {
	var x reflect.SliceHeader
	x.Len = sizeOfMyStruct
	x.Cap = sizeOfMyStruct
	x.Data = uintptr(unsafe.Pointer(s))
	return *(*[]byte)(unsafe.Pointer(&x))
}

func BytesToMyStruct(b []byte) *MyStruct {
	return (*MyStruct)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&b)).Data))
}

func main() {
	var myMyStruct = MyStruct{"bydsfssf"}
	var by = MyStructToBytes(&myMyStruct)
	fmt.Println(by)
	var my = BytesToMyStruct(by)
	fmt.Println(my)
}
