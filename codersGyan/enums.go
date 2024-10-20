package main

import "fmt"

// type OrderStatus int
type OrderStatus string

const (
	// Received OrderStatus = iota  // int
	Received  OrderStatus = "Received"
	Confirmed             = "Confirmed"
	Prepared              = "Prepared"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println(status)
}

func main() {
	changeOrderStatus(Received)
}
