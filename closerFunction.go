package main

import "fmt"

func company() func() int {

	a := 0

	return func() int {
		a++
		return a
	}

}

func main() {

	data := company()

	fmt.Println(data())
	fmt.Println(data())
	fmt.Println(data())

}
