package main

import "fmt"

func main() {
	a := 23
	b := 34
	c := 2
	variadicFunc(a, b, c)

	// anonimus function
	func() {
		fmt.Println("TEsting")
	}()

	// anonimus function with parameter
	func(s ...string) {
		fmt.Printf("TEsting first value : %v \n second value: %v \n ", s[0], s[1])
	}("wap", "institute")

	// put function in variable
	fu := func() int {
		return 108
	}

	fmt.Printf("return value is %v \n", fu())

}

func variadicFunc(a ...int) {
	fmt.Println(a)
}

func init() {
	fmt.Println("init functioin will not take arguments and not return and value and it' will call first")
}
