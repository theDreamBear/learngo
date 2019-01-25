package main

// case 只会运行一个
func testSwitch(i int) {
	switch {
	case i > 0:
		{
			println("positive number")
		}
	case i > -4:
		{
			println("near zero")
		}
	default:
		{
			println("error")
		}
	}
}

func testRange(){
	 x := []int{10,20,30} // int[]{1,2,3}  和c++比相反
	for i, n := range x{
		println(i, " : ", n)
	}
}

func main() {
	//testSwitch(1)
	testRange()
}