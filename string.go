package main

import (
	"fmt"
	"strings"
)

func main() {

	// trim strings and space

	a := " @testing@ "
	a = strings.TrimSpace(a)
	a = strings.TrimLeft(a, "@")
	fmt.Println(a)
	a = strings.Trim(a, "@")
	fmt.Println(a)

	// trim suffix

	b := "hello world"
	b = strings.TrimSuffix(b, "world")
	fmt.Println(b)

	// trim preffix

	c := "hello world"
	c = strings.TrimPrefix(c, "hello")

	fmt.Println(c)

	for index, val := range a {
		fmt.Println(index, string(val))
	}
	fmt.Println("type assertion ", []byte(a))

}
