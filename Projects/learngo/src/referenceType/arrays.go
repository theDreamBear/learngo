package main

import "fmt"

/*
*  数组大小也是数组的一部分
*  数组是值类型的， 	传递的时候会发生拷贝
 */

func arrayTransAsValue(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func arrayTransAsPreference(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func arrayInit() {
	var (
		arr1 [3]int
		arr2 = [4]int{1, 2, 43}
		arr3 = [...]int{111}
	)
	fmt.Println(arr1, arr2, arr3)
}

func arrayForeach() {
	var arr = [...]int{0, 10, 20, 30}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}

}

func main() {
	arrayInit()

	arrayForeach()

	arr := [5]int{0, 10, 20, 30, 40}
	arrayTransAsValue(arr)
	arrayTransAsPreference(&arr)
	fmt.Println("after cahnge")
	fmt.Println(arr)
}
