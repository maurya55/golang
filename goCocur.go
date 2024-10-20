package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go count("testing 1")

	go count("testing 2")

	fmt.Println("End program")
	wg.Wait()

}

func count(val string) {

	for i := 0; i < 5; i++ {
		fmt.Println(val)

	}
	wg.Done()

}
