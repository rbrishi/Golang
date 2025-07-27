package main

import "fmt"

//enums are enumerated types that are used to define a set of named constants

type orderStatus int
const(
	Received orderStatus = iota
	Confirmed
	Prepared
	Delivered

	// Received orderStatus = "Received"
	// Confirmed = "Confirmed"
	// Prepared = "Prepared"
	// Delivered = "Delivered"
)

func changeOrderStatus(status orderStatus){
	fmt.Println("Order status changed to:", status)
}


func main(){
	changeOrderStatus(Received)
	changeOrderStatus(Confirmed)
	changeOrderStatus(Prepared)
	changeOrderStatus(Delivered)
}