package main

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

func evaluation(a, b int, op string) (result int,err  error){
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b , nil
	case "*":
		return  a * b , nil
	case "/":
		result, _, err := divideWithMod(a, b)  //性能不好
		return result, err
	default:
		return 0, errors.New("error: undefined op")
	}
}

func divideWithMod(a, b int)(int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("error: divide zero")
	} else {
		return a / b, a % b, nil
	}

}

// 下面的实现有不够好
func addF(a , b int)(int, float64, error) {
	return a + b, 0, nil
}

func subF(a, b int)(int, float64, error) {
	return a - b , 0, nil
}

func multiplyF(a, b int)(int, float64, error) {
	return a * b, 0, nil
}

func divideF(a, b int)(int, float64, error) {
	if b == 0 {
		return 0, 0, errors.New("error: divide zero")
	} else {
		var c float64
		if a > b {
			c = float64(a - b)/ float64(b)
		} else {
			c = float64(a) / float64(b)
		}
		return a / b, c , nil
	}
}

// 函数式编程
func apply(operator func(int, int) (int, float64, error),a, b int )(int, float64,error) {
	p := reflect.ValueOf(operator).Pointer()
	name := runtime.FuncForPC(p).Name()
	fmt.Printf("op is %s\n ", name)
	if q, r, err := operator(a, b); err != nil{
		return 0, 0, err
	} else {
		return q, r, nil
	}
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func swap1(a, b int)(int, int)  {
	return b, a
}

func main() {
	//if result, err := evaluation(12, 0,"/"); err != nil{
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(result)
	//}
	//
	//q, r, _ := divideWithMod(4,3)
	//fmt.Printf("%d / %d answer is %d, mod is %d", 4, 3, q, r)

	a, b := 3, 4
	if i, i2, e := apply(divideF, a, b); e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("%f",float64(i) + i2)
	}
}
