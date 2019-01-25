package main

import (
	"fmt"
)


func commonVar(){

	// S1
	var i int
	fmt.Println("i 的指是： ",i)
	var y = false
	fmt.Println("y == false ", y ==false)

	// S2
	var name, age = "nic", 18
	fmt.Println("my name is:", name, "\t my age is : ", age)

	// S3
	var(
		newName= "TT"
		grade, newAge int
		sex string
	)
	fmt.Println("my new name is ", newName, "\t I am in grade ", grade, "\t my born age is ", newAge, "\t my sex is ", sex)

	// S4
	x := 100
	fmt.Println("x is ", x)
}

var x = 100
func varLocalAndAll() {
	println(&x , x)

	// 下面模式
	// 只能用在函数内部
	// 定义并显示声明
	// 不能提供类型
	x := "abc"
	println(&x, x)
}


func shorrModeDepression() {
	// 最少有一个新变量被定义
	// 必须相同定义域
	// S1
	v := 100
	v, n := 200, 300
	println(v, n)

	//m := 100
	// 没有新的变量
	//m := 200
	println(&x, x)
	x, n := 200, 3
	println(&x ,x)
	println(n)

}

//multi param=
func multiparam(){
	x, y := 1, 2
	x, y = y + 3, x + 2
	println("x = ",x)
	println("y = ", y)
}

func main() {
	//commonVar()
	//varLocalAndAll()
	//shorrModeDepression()
	//multiparam()
	//var 你擦后 = "nic"
	//println(你擦后)
}
