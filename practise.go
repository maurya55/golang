package main

import "fmt"

type Email struct {
	email1 string
	email2 string
}

type Employee struct {
	name  string
	email *Email
}

func main() {

	var val *int
	var num int
	val = &num
	num = 33
	fmt.Println(*val)

	em := &Email{email1: "wap@", email2: "ins@"}

	result := Employee{name: "Wapins", email: em}

	fmt.Printf("%+v", result)

}
