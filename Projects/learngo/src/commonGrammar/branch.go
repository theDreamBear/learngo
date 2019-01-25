package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func ReadFiles(path string) {
	if content, err := ioutil.ReadFile(path); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}
}



func eval(lhs , rhs int, op string) (int, error){
	ret := 0
	var err error
	switch op {
	case "+":
		ret =  lhs + rhs
	case "-":
		ret = lhs - rhs
	case "*":
		ret = lhs * rhs
	case "/":
		if rhs == 0{
			err = errors.New("除以0错误")
		} else {
			ret = lhs / rhs
		}
	default:
		err = errors.New("未定义运算浮")
	}
	return ret, err
}

func grade(score int)(string, error) {
	var(
		g string
		err error
	)
	switch{
	case score < 0 || score > 100 :
		err = errors.New("score is wrong")
		panic(fmt.Sprintf("%s",err))
	case score < 60:
		g = "不及格"
	case score < 80:
		g = "中等"
	case score < 90:
		g = "良"
	case score < 100:
		g = "优秀"
	default:
		g = "superlaxi"
	}
	return g, err
}

// fmt.Println(....), 所有都计算再显示
func caculateSequence(i int) int {
	switch {
	case  i < 10:
		panic("wrong")
	default:
		return i + 1

	}
	return 100
}



func main() {
	const file = "./abc.txt"
	//ReadFiles(file)
	if ret, err := eval(10, 0, "abs"); err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(ret)
	}

	if g, err := grade(99); err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(g)
	}

}
