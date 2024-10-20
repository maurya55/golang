package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("kkk")
	wg.Add(1)
	ret1 := f("goroutine", &wg)
	fmt.Println(ret1)
	wg.Wait()

	wg.Add(2)
	ret2 := f("Testing 1 ", &wg)
	fmt.Println(ret2)

	ret3 := f("Testing 2", &wg)
	fmt.Println(ret3)

	wg.Wait()
	fmt.Println("==========>")

	// var wg sync.WaitGroup
	// fmt.Println("kkk")
	// wg.Add(1)
	// go f("goroutine", &wg)
	// wg.Wait()
	// wg.Add(2)
	// go f("Testing 1", &wg)
	// go f("Testing 2", &wg)
	// wg.Wait()
	// fmt.Println("==========>")
}

// func d(dem string, wg *sync.WaitGroup)
// {

// }

func f(from string, wg *sync.WaitGroup) string {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i, ":", from)
		time.Sleep(time.Second * 1)

	}

	return from

}
