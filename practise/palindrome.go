package main

import "fmt"

func main() {

	st := "1234321"

	i := 0
	checkPalindrom := false
	for i < len(st) {
		if st[i] == st[len(st)-1-i] {
			checkPalindrom = true
		} else {
			checkPalindrom = false
			break
		}
		i++

	}

	if checkPalindrom {
		fmt.Printf("palindrom string %v", st)
	} else {
		fmt.Printf("not palindrom string %v", st)

	}

}
