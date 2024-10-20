package main

import "fmt"

type Employee interface {
	PrintUser(name string)
}

type Emp int

func (id Emp) PrintUser(name string) {
	fmt.Printf("employee id is %v and employee name is %v\n", id, name)
}

func main() {
	var e1 Employee
	e1 = Emp(12)
	e1.PrintUser("dfa")
	e1 = Emp(13)
	e1.PrintUser("falllaa")
}
