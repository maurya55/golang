package main

import "fmt"

type Employee struct {
	name string
	age  int
}

// func (emp *Employee) add() {
func (emp Employee) add() {
	// *&(emp).age = 14
	emp.age = 14

}

func main() {

	var a = []int{1, 2, 3, 4}

	a = append(a[:1], a[3:]...)

	fmt.Println(a)

	var em = Employee{
		name: "Wap institute",
		age:  12,
	}

	em.add()

	fmt.Println(em)

}
