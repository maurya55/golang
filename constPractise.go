package main

import "fmt"

const (
	first = iota
	second
	third = "a"
)

func main() {

	a := 10
	b := 20
	c := 30

	fmt.Println(first, second, third)

	data, _ := Test(a, b, c)
	fmt.Println(data)

}

func Test(num1 int, num2 int, num3 int) (int, error) {
	return num1 + num2 + num3, nil
}
