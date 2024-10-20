package main

import "fmt"

func deleteSpecificPosition(position int, data []int) []int {

	data = append(data[:position], data[position+1:]...)

	return data
}

func main() {

	var arrayData = []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(arrayData)

	fmt.Println("delete 2 position", deleteSpecificPosition(2, arrayData))
	fmt.Println("delete 4 position", deleteSpecificPosition(4, arrayData))
	fmt.Println("delete 5 position", deleteSpecificPosition(5, arrayData))

}
