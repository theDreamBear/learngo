package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `name`
	Age  int    `age`
}

func main() {
	var u User
	u.Name = "nil"
	u.Age = 12
	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		fmt.Println(sf.Tag)
	}

	var rv = reflect.ValueOf(&u)
	var rr = rv.Interface().(*User)
	rr.Name = "lsllss"
	fmt.Println(reflect.ValueOf(u).Interface())
	fmt.Println(rr)
}
