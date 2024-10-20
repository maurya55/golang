package main

import "fmt"

func main() {

	num := 9
	var A int
	var B int

	for i := 0; i < num; i++ {
		// fmt.Println("test")
		if i == 0 {
			A = 0
			B = 1
		} else {
			c := B
			B = A + B
			A = c

		}

		fmt.Println(B)
	}

	fmt.Println(B)
}
