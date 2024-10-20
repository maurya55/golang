package main

import "fmt"

type Employee struct {
	Salary1 int
	Salary2 int
}

func main() {
	fmt.Println("Testing")

	count := Employee{Salary1: 12, Salary2: 12}
	count.add()

	// anonmus struct

	test := struct {
		name  string
		email string
	}{
		"wap",
		"wap@gmail.com",
	}

	fmt.Println(test.email)

}

func (e Employee) add() {
	fmt.Println(e.Salary1 + e.Salary2)
}
