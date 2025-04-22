package main

import (
	"fmt"
	"time"
)

type order struct {
	id string
	amount float32
	status string
	createdAt time.Time //nano sec precision
}

func newOrder(id string, amount float32, status string, createdAt time.Time) *order {
	// makeshift constructor
	myOrder := order {
		id: id,
		amount: amount,
		status: status,
		createdAt: createdAt,
	}
	return &myOrder
}

//receiver type implementation of Method of class order
func (o *order) changeStatus(status string) { //pointer ;)
	o.status = status
}

func main() {
	// classes of GO
	myorder := order{
		id: "1",
		amount: 99.99,
		status: "pending",
		createdAt: time.Now(),
	}
	//instance or object
	fmt.Println(myorder)
	myorder.changeStatus("delivered")
	fmt.Println(myorder)

	conOrder := newOrder("1", 9.9, "confirmed", time.Now())
	fmt.Println(conOrder)

	language := struct {
		name string
		isGood bool
	} {"golang", true}
	fmt.Println(language)

	//embedding of structs is also possible
}