package main

import "fmt"

type Employee struct {
	Name string
	age  int
}

func (e *Employee) changeName(newName string) {

	(*e).Name = "Wap institute"

}

func changeName(st *string) {

	*(st) = "institute"
	// fmt.Println(*st)
}

func main() {

	var data string = "wap"

	fmt.Println(data)
	changeName(&data)
	fmt.Println(data)

	var employeeData = Employee{
		Name: "Wapinstitute",
		age:  25,
	}

	fmt.Printf(" before name change %+v\n", employeeData)

	employeeData.changeName("Wap institute")

	fmt.Printf(" after name change %+v\n", employeeData)

}
