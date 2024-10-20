package main

import "fmt"

func main() {

	var data = []int{5, 9, 13, 17, 45, 67, 89, 100}

	var find = 91

	val, err := binerySearch(find, data)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("value found in %v position", val)

}

func binerySearch(find int, data []int) (int, error) {

	var start = 0
	var end = len(data) - 1
	var position int
	var checkBool = false
	for end >= start {

		mid := (start + end) / 2
		// var data = []int{5, 9, 13, 17, 45, 67, 89, 100}

		// fmt.Println(start, end, mid)
		if data[mid] == find {
			position = mid
			checkBool = true
			break
		} else if data[mid] > find {
			end = mid - 1
		} else {
			start = mid + 1
		}

	}

	if !checkBool {
		return position, fmt.Errorf("%v value not found", find)
	}

	return position, nil

}
