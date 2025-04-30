package main

import "fmt"


type paymenter interface{
	pay(amount float32)
}
// interface is a contract where we define accepted methods of a struct which is accepted.
// open close principle

type payment struct {
	gateway paymenter
}


func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct {

}
type stripe struct {

}

func (r razorpay) pay(amount float32) {
	fmt.Println(amount)
}

func (s stripe) pay(amount float32) {
	fmt.Println("stripe ",amount)
}

func main() {
	stripePaymentGateway := stripe{}
	// razorPayPaymentGateway := razorpay{}
	paymentAmount := payment{
		gateway: stripePaymentGateway,
	}
	paymentAmount.makePayment(9.99)
}