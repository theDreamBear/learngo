package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `name`
	Age  int    `age`
}

func (User) DoSomething(name string, age int) {
	fmt.Println("today is a good day")
}

func DoFiledAndMethod(input interface{}) {
	var ttype = reflect.TypeOf(input)
	var tvalue = reflect.ValueOf(input)

	for i := 0; i < ttype.NumField(); i++ {
		field := ttype.Field(i)
		vv := tvalue.Field(i).Interface()
		fmt.Printf("%d:  %s: %v = %v\n", i, field.Name, field.Type, vv)
	}

	for i := 0; i < ttype.NumMethod(); i++ {
		m := ttype.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
		med := tvalue.MethodByName(m.Name)
		//var s = tvalue.Field(0).Interface().(string)
		//var a = tvalue.Field(1).Interface().(int)
		args := []reflect.Value{reflect.ValueOf(tvalue.Field(0).Interface().(string)), reflect.ValueOf(tvalue.Field(1).Interface().(int))}
		med.Call(args)
	}
}

func main() {
	var user User
	user.Name = "heecfsdf"
	user.Age = 18
	DoFiledAndMethod(user)

	var a = 10
	var vi = reflect.ValueOf(&a)
	pt := vi.Elem()
	if pt.CanSet() {
		pt.SetInt(10)
	}
	fmt.Println(a)
}
