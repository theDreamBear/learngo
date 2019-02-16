package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func Minit() {

	var Rv = reflect.ValueOf(nil)
	if Rv.IsValid() {
		fmt.Println("zero")
	} else {
		fmt.Println("none zero")
	}

}

func getReflect(i interface{}) {
	var x = reflect.TypeOf(i)
	var y = reflect.ValueOf(i).Interface()
	fmt.Printf("%T: %v\n", y, y)
	fmt.Println(x)
}

func getField(i interface{}) {
	Type := reflect.TypeOf(i)
	Value := reflect.ValueOf(i)

	fmt.Printf("the type of i is %T, value is %v\n", Type, Type)
	fmt.Printf("the type of i is %T, value is %v\n", Value, Value)

	for i := 0; i < Type.NumField(); i++ {
		fd := Type.Field(i)
		name := fd.Tag.Get("nic")
		fmt.Printf("%T, %v\n", name, name)
		ve := Value.Field(i).Interface()
		fmt.Printf("%v\n", fd)
		fmt.Printf("%s: %v = %v\n", fd.Name, fd.Type, ve)
	}

	for i := 0; i < Type.NumMethod(); i++ {
		m := Type.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

type User struct {
	Name string `nic:"name"`
	Age  int    `nic:"age"`
	Sex  string `nic:"sex"`
}

func (user User) String() string {
	return fmt.Sprintln("%s", user.Name)
}

func ImproveStructReflectPerformance() {
	slice1 := []string{"hello", "world"}
	//type sliceHeader struct {
	//	Data unsafe.Pointer
	//	Len  int
	//	Cap  int
	//}
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice1))
	fmt.Println(header.Len)
	elementType := reflect.TypeOf(slice1).Elem()
	secondElementPtr := uintptr(header.Data) + elementType.Size()
	*((*string)(unsafe.Pointer(secondElementPtr))) = "!!!"
	fmt.Println(slice1)
}

func main() {
	//Minit()
	//i := 100
	//getReflect(i)

	//var user = User{"nic", 20, "man"}
	//getField(user)

	ImproveStructReflectPerformance()
}
