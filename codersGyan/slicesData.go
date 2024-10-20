package main

import "fmt"

func main() {

	var num = []int{1, 2, 3}
	num = append(num, 4, 5)
	fmt.Println(num)

	var num1 = make([]int, 2, 3)
	num1 = append(num1, 34)
	fmt.Print(num1)

}
