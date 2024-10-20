package main

import (
	"fmt"
)

func arraySearch(userVal int, arrayData []int) (int, error) {

	var num int
	var checkBool = false

	for ind, val := range arrayData {
		if val == userVal {
			num = ind
			checkBool = true
			break

		}
	}
	if !checkBool {
		return 0, fmt.Errorf("%v is not found", userVal)
	}

	return num, nil
}

func main() {

	var arrayData = []int{20, 40, 60, 5, 10, 70, 80, 99}

	val, err := arraySearch(40, arrayData)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(" 20 is found on position of %v \n", val)

}
