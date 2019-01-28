package main

import (
	"errorhandling/err"
	"fmt"
)

type MserError interface {
	err.UserError
}

type McError struct {
}

func (mc McError) Message() string {
	return fmt.Sprintln("world world")
}

func main() {
	var mc McError
	fmt.Println(mc.Message())
}
