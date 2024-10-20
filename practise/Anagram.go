package main

import "fmt"

func main() {

	fmt.Println(checkAnagram("silent", "listen"))

}

func checkAnagram(str1 string, str2 string) bool {

	var checkAng = true
	if len(str1) != len(str2) {
		checkAng = false
	}

	var ab = make(map[string]int)

	for _, val := range str1 {
		ab[string(val)] = ab[string(val)] + 1
	}

	for _, val := range str2 {
		ab[string(val)]--
	}

	for _, val := range str2 {
		if ab[string(val)] != 0 {
			checkAng = false
			break
		}
	}

	// fmt.Println(ab)
	return checkAng
}
