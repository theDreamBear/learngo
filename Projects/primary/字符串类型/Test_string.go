package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

//string
/*   type stringStruct struct{
*		str unsafe.Pointer
*		len int
*	}
*   空字符 	！= nil,  == "" len == 0
*   注意：其指针也不为空
**
*/

func str_test(){
	var s string
	if s == "" {
		fmt.Println("string is empty")
	}
	if (*reflect.StringHeader)(unsafe.Pointer(&s)) != nil {
		fmt.Println("string unsafe Pointer not to nil")
		fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s)))
	}
}

// `` 	包裹的字符串 raw 字符串 可以多行
func str_test2() {
	var s =
		`
早上我起床
实在起不来
到底为什么
不明白！
           `
	fmt.Println(s)

	path := `c:\home\homework\dd\`
	fmt.Println(path)
}

// 可以按引用返回字符串数组， 但不能获取元素地址
func str_test3() {
	s := "abc"
	println(s[0])
	//println(&s[1]) //不能获取 数组元素地址
}

// 以切片语法返回字串时， 其内部仍然指向原字节数组

func str_slice_test() {
	s := "abcdefg"
	s1 := s[:]
	fmt.Printf("%#v\n",(*reflect.StringHeader)(unsafe.Pointer(&s)))
	fmt.Printf("%#v\n",(*reflect.StringHeader)(unsafe.Pointer(&s1)))
}

// 字符串遍历
func str_foreach() {
	s := "abcdefg弟弟"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d: [%c]\n", i, s[i])
	}

	for i, c := range s {
		fmt.Printf("%d: [%c]\n", i, c)
	}
}

// 修改字符串的值
func str_change(s string, ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))

	fmt.Printf(s, *h)
}

func main() {
	//str_test()
	//str_test2()
	//str_slice_test()
	//str_foreach()
	//s := "hello world"
	//str_change("s:%x\n", &s)
	r := 'h'
	println(utf8.RuneCountInString(r))
}
