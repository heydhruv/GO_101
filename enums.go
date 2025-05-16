package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Shipped
	Delivered
)

func changeOrderStatus(status OrderStatus){
	fmt.Println("Changed order status to", status)
}

func main() {
	changeOrderStatus(Received)
}