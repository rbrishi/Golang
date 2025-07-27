package main

import "fmt"

type PaymentGateway interface {
	pay(amount float64)
}

type Payment struct {
gateway PaymentGateway
}

func (P Payment) makePayment(amount float64) {
	P.gateway.pay(amount)
}


type Razorpay struct {

}


func (R Razorpay) pay(amaount float64) {
	fmt.Println("making payment of amount:", amaount)
}
func main(){
	newPayment := Payment{}
	newPayment.makePayment(100.50)
}
