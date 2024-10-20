package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 9; i++ {
			fmt.Println(" first log : ", i)
		}
	}()
	// wg.Wait()

	// wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 9; i++ {
			fmt.Println(" second log : ", i)
		}
	}()

	wg.Wait()

}
