package main

import "fmt"

type Data struct {
	val   int
	left  *Data
	right *Data
}

func main() {

	root := &Data{}

	addData(root, 100)
	addData(root, 95)
	addData(root, 80)
	addData(root, 101)

	// fmt.Println(root)
	// printData(root)

}

// func printData(da *Data) {
// 	if da.val != 0 {
// 		fmt.Println(da.val)
// 		if da.right != nil {
// 			printData(da.right)
// 		} else if da.right != nil {
// 			printData(da.left)
// 		}
// 	}
// }

func addData(da *Data, val int) {

	if da.val == 0 {

		da.val = val
		fmt.Println(da)
	} else {
		insertData(da, val)
	}
}

func insertData(da *Data, val int) {
	fmt.Println(val)
	if val > da.val {
		if da.right == nil {
			newData := Data{val: val}
			da.right = &newData
		} else {
			insertData(da.right, val)
		}
	} else {
		if da.left == nil {
			newData := Data{val: val}
			da.left = &newData
		} else {
			insertData(da.left, val)
		}
	}
}
