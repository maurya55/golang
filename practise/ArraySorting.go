package main

import "fmt"

func main() {

	var arrayData = []int{34, 3, 1, 25, 54, 12, 89, 5, 76, 789, 1}

	for i, _ := range arrayData {
		fmt.Print(i)
		for j, jVal := range arrayData {
			if j < len(arrayData)-1 {
				if jVal > arrayData[j+1] {
					var a = arrayData[j+1]
					arrayData[j+1] = jVal
					arrayData[j] = a
				}
			}
		}
	}

	fmt.Println(arrayData)

	arrayData = []int{34, 3, 1, 25, 54, 12, 89, 5, 76, 789, 1}

}
