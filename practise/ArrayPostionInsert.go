package main

import "fmt"

func ArrayPostionInsert(position int, num int, arrayData []int) []int {

	arrayData = append(arrayData[:position], append([]int{num}, arrayData[position:]...)...)

	return arrayData

}

func main() {

	var arrayData = []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(arrayData)

	fmt.Println("Add 2 position 22", ArrayPostionInsert(2, 22, arrayData))
	fmt.Println("Add 2 position 22", ArrayPostionInsert(4, 44, arrayData))
	fmt.Println("Add 2 position 22", ArrayPostionInsert(5, 55, arrayData))

}
