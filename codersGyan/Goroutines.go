package main

import "time"

func task(d string) {

	for i := 0; i < 5; i++ {
		println(d, i)
	}

}

func main() {

	go task("first")
	go task("second")

	time.Sleep(time.Second * 1)

}
