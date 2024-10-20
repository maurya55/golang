package main

import "fmt"

func counter() func() int {

	return func() int {
		return 33
	}

}

func main() {

	ck := counter()

	fmt.Println(ck())

}
