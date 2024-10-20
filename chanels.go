package main

import (
	"fmt"
	"time"
)

func main() {

	message := make(chan []int, 3)

	//  sync.WaitGroup
	go demo("testing 1", message)

	// go demo1(message)

	// x := <-message
	// fmt.Println(x)
	for i := range message {
		fmt.Println("print chanel : ", i)
	}
	time.Sleep(time.Second * 3)

	fmt.Println("End programm")

}

func demo1(ch chan int) {

	for i := range ch {
		fmt.Println(i)
	}

}

func demo(val string, ch chan []int) {
	var ab = []int{22, 33, 22, 11}
	// fmt.Println("loop ")
	// for i := 11; i < 20; i++ {
	// 	fmt.Println("loop ", 1)
	// 	ch <- i
	// }
	ch <- ab
	// time.Sleep(time.Second * 3)
	close(ch)
	println("demo function", val)

}
