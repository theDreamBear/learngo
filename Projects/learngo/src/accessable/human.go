package accessable

import "fmt"

type Hunman struct {
	man person
}

func (hn Hunman) Eat() {
	//fmt.Printf("name:%s\tage:%d\tsex:%s\n", man.name, man.age, man.sex)
	fmt.Println("hello world")
}
