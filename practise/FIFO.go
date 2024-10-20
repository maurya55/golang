package main

import "fmt"

func main() {

	var queue []int

	queue = enqueue(10, queue)
	fmt.Println("After pushing 10 ", queue)

	queue = enqueue(20, queue)
	fmt.Println("After pushing 20 ", queue)

	queue = enqueue(30, queue)
	fmt.Println("After pushing 30 ", queue)

	queue = enqueue(40, queue)
	fmt.Println("After pushing 40 ", queue)

	queue = enqueue(50, queue)
	fmt.Println("After pushing 50 ", queue)

	queue = dequeue(queue)
	fmt.Println("After removing first value ", queue)

	queue = dequeue(queue)
	fmt.Println("After removing first value ", queue)
}

func enqueue(num int, queue []int) []int {

	queue = append([]int{num}, queue...)

	// var da = []int{num}
	// for _, val := range queue {
	// 	da = append(da, val)
	// }
	return queue

}

func dequeue(queue []int) []int {

	if len(queue) == 1 || len(queue) == 0 {
		queue = []int{}
	} else {
		queue = queue[1:]
	}

	return queue

}
