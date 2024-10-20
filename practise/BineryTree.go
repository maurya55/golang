package main

import "fmt"

type Child struct {
	name  string
	child *Child
}

// type Parent struct {

// }

func main() {

	root := Child{name: "wap"}

	t1 := Child{name: "Institute"}
	t2 := Child{name: "abc"}
	t3 := Child{name: "xyz"}
	t4 := Child{name: "def"}

	root.child = &t1
	t1.child = &t2
	t2.child = &t3
	t3.child = &t4

	printData(&root)

}

func printData(data *Child) {

	if data == nil {
		return
	}

	fmt.Println(data.name)
	printData(data.child)

}
