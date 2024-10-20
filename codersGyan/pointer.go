package main

import "fmt"

func changeName(num *int) {
	*num = 10
}

func main() {

	num := 12

	fmt.Println(num)

	changeName(&num)

	fmt.Println(num)

}
