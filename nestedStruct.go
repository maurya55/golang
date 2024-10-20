package main

import "fmt"

type AddressStruct struct {
	Address1 string
	Address2 string
}
type Employee struct {
	Name    string
	Email   string
	Address []AddressStruct
}

func main() {

	result := Employee{
		Name:  "Wap",
		Email: "wap@gmail.com",
		Address: []AddressStruct{
			{Address1: "sion", Address2: "chunabhatti"},
			{Address1: "sion", Address2: "GTB"},
		},
	}

	fmt.Println(result)

}
