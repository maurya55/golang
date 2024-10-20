package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// fmt.Println()
	fmt.Println(checkAnagram("anil", "ilan"))
}

func checkAnagram(str1 string, str2 string) bool {

	st1 := strings.Split(str1, "")
	st2 := strings.Split(str2, "")
	// ad := slices.Sort(st1)

	// var st3 = make([]string, len(st1))
	// copy(st3, st1)

	sort.Strings(st1)
	sort.Strings(st2)

	str1 = strings.Join(st1, "")
	str2 = strings.Join(st2, "")

	// fmt.Println(st1, st2, st3)
	fmt.Println(str1, str2)
	return str1 == str2

	// fmt.Println(reflect.ValueOf(st1).Kind() == reflect.Array)

}
