package main

import "fmt"

type Person struct {
	Name  string
	Email string
}

// func (pe *Person) changeEmail(em string) {
// 	(*pe).Email = em

// }

func main() {

	// setValue := Person{Name: "Wap", Email: "Wap@gmal.com"}

	// setValue.changeEmail("institute@gmail.com")

	// fmt.Println(setValue)

	var value int
	var ptr *int
	ptr = &value
	fmt.Println("ptr value is : ", ptr)
	fmt.Println("value is  : ", value)
	*ptr = 12
	fmt.Println("changed value thorough ptr ", value)
	fmt.Println("ptr value ", ptr, *ptr)

}
func init() {
	fmt.Println("============>")
}
