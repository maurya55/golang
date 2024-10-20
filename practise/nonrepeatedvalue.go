package main

import "fmt"

func main() {

	input := "aprogrammingp"

	// fmt.Println("No non-repeated character found.")

	var da = make(map[string]int)

	for _, val := range input {
		da[string(val)] = da[string(val)] + 1
	}

	for _, val := range input {
		if da[string(val)] == 1 {
			fmt.Printf("first value found %v \n", string(val))
			return
		}
	}

	fmt.Println("non repeated value not found ")

}
