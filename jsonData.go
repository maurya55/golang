package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Human struct {

		// defining struct variables
		Name string
		Age  int
	}

	human1 := Human{"Ankit", 23}

	human_enc, err := json.Marshal(human1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(human_enc))

}
