package main

import "fmt"

func main() {

	var a *int
	var b *int

	a = ptr(19)
	b = ptr(19)
	c := *a + *b
	fmt.Println(a, b, c)

	z := ptr(c)
	fmt.Println(*z)

	var member *int = new(int)
	var val int
	val = 12
	member = &val

	fmt.Println(member)

}

func ptr(v int) *int {
	return &v
}
func ptr1[T any](v T) *T {
	return &v
}
