package main

import (
	"fmt"
)

func main() {

	var map1 = make(map[int]string)
	map1[11] = "wap"
	map1[12] = "institute"
	fmt.Println("map1======> ", map1)

	var map2 = map[int]string{
		22: "first",
		23: "second",
		24: "four",
	}

	fmt.Println("map2========> ", map2)

	fmt.Println("nested map=============>")

	var nestedMap = make(map[string]map[string]int)
	nestedMap["one"] = make(map[string]int)

	nestedMap["one"]["two"] = 12

	fmt.Println(nestedMap, nestedMap["one"]["two"])

}
