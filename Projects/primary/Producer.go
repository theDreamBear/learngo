package main

func product(data chan int) {
	for i:= 0 ; i < 10; i++ {
		data <- i;
	}
	close (data)
}


func consumer(data chan int, done chan bool) {
	for x := range data{
		println("recv: ", x)
	}
	done <- true
}

func main () {
	done := make (chan bool)
	data := make (chan int,100)

	go product(data)
	//time.Sleep(time.Microsecond * 2)
	go consumer(data, done)
	//`time.Sleep(time.Microsecond)
	<-done
}