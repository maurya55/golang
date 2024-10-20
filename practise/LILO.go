package main

import "fmt"

var queue []int

func main() {

	// var queue = make([]int, 0)

	queue = enqueue(10)
	fmt.Println("After pushing 10 ", queue)

	queue = enqueue(20)
	fmt.Println("After pushing 20 ", queue)

	queue = enqueue(30)
	fmt.Println("After pushing 30 ", queue)

	queue = enqueue(40)
	fmt.Println("After pushing 40 ", queue)

	queue = enqueue(50)
	fmt.Println("After pushing 50 ", queue)

	queue = dequeue()
	fmt.Println("After removing last value ", queue)

	queue = dequeue()
	fmt.Println("After removing last value ", queue)

	queue = enqueue(100)
	fmt.Println("After pushing 100 ", queue)

}

func enqueue(a int) []int {

	queue = append(queue, a)
	return queue

}

func dequeue() []int {

	if len(queue) == 1 || len(queue) == 0 {
		queue = []int{}
	} else {
		queue = queue[:len(queue)-1]
	}

	return queue

}
