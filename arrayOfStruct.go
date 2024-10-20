package main

import "fmt"

type Employee struct {
	name  string
	email string
	age   int
}

func main() {

	// result := make([]Employee, 1, 5)
	result := make([]Employee, 0, 0)

	fmt.Println(result)

	result = append(result, Employee{name: "wap", email: "wap@gmail.com", age: 12})

	fmt.Println(result)

}
