package main

import "fmt"

func main() {

	var ar [3]int

	ar[0] = 32
	ar[2] = 36

	for ind, value := range ar {
		fmt.Println(ind, value)
	}

	var ar1 = ar
	ar1[1] = 12

	fmt.Println(ar1)

	a := [5]int{1, 2, 3, 4, 5}

	sl := make([]int, 4)
	copy(sl, a[0:4])

	sl[0] = 11

	fmt.Printf("array data is %v\n", a)
	fmt.Printf("slice data is %v\n", sl)

}
