package main

import "fmt"

type User struct{
	name string
	age int
	passwd string
}

func simulateLogging() {
	var user User
	_, ok := fmt.Scanln(&user.name, &user.age, &user.passwd)
	if ok == nil {
		if user.name == "nic" && user.age == 18 && user.passwd == "123456" {
			println("logging success")
		} else {
			println("false logging please retry")
			simulateLogging()
		}
	} else {
		println("input formate is wrong, please try again")
		return
	}
}

func main () {
	simulateLogging()
}