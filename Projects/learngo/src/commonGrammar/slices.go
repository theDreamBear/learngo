package main

import "fmt"

// 切片前开后闭
func sliceInit(){
	arr := [...]int{0,1,2,3,4,5,6,7,8,9}
	s := arr[2:6]
	fmt.Println(s)
}

func updateSlice(s [] int){
	s[0] = 100
}




func sliceslice() {
	sce := []int{1,2,3,4,5,6}
	printSlice(sce)
	s := sce[1:5]
	printSlice(s)
}

func main() {
	//sliceInit()
	//arr := [...]int{0,1,2,3,4,5,6,7,8,9}
	//s := arr[1:7]
	//fmt.Println(s)
	//updateSlice(s)
	//fmt.Println(s)
	sliceslice()
}
