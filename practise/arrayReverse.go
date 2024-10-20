package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("testing")

	var st = "wap"

	da := strings.Split(st, "")
	var final []string
	for i := 0; i < len(da); i++ {
		j := len(da) - 1 - i
		final = append(final, da[j])
	}

	fmt.Println(strings.Join(final, ""))
}
