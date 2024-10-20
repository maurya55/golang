package main

import "fmt"

func main() {
	fmt.Println("one")
	defer fmt.Println("Two")
	fmt.Println("Three")
	defer Demo()
}

func Demo() {
	fmt.Println("Demo function data")
}
