package main

import (
	"fmt"
)

func main() {

	arr := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	fmt.Printf("2D array data : %v\n", arr)
	arr[1][1] = 66
	fmt.Printf("2D array data : %v\n", arr)

	delId := 1
	originalarr := []string{"Vatsal", "Kartik", "Ritika", "Veronica", "Isha"} //create array
	fmt.Println(originalarr[delId:])

	fmt.Println("============>", append(originalarr[:delId], originalarr[delId+1:]...))

	fmt.Println("The original array is:", originalarr)

	erroMessage, newArr, err := remove_ele(originalarr[:], 3)
	if err != nil {
		fmt.Print(err)
	} else if erroMessage != "" {
		fmt.Printf("error is : %v\n", erroMessage)
	} else {
		fmt.Printf("final value is %v", newArr)
	}

}

func remove_ele(originalarr []string, deleteId int) (string, []string, error) {

	if deleteId < 0 || deleteId > len(originalarr) {
		return "deletedId is out fo bound", nil, fmt.Errorf("this is error")
	}

	j := 0

	for index, _ := range originalarr {
		// fmt.Println(index, key)
		if index != deleteId {
			originalarr[j] = originalarr[index]
			j++
		}
	}

	newarr := originalarr[:j]

	// time.Sleep(time.Second * 2)

	return "", newarr, nil

}
