package main

import (
	"fmt"
	"time"
)

// Structs in Go are used to create custom data types by grouping related fields together
// They are similar to classes in other languages but without inheritance

// Orders struct represents an order in a system
// It demonstrates both basic field types and struct composition
type Orders struct {
	ID        int
	Name      string
	Price     float64
	CreatedAt time.Time  // Using time.Time for timestamp with nanosecond precision
	Customer  Customer   // Embedding another struct - demonstrates composition
}

// Customer struct represents a customer in the system
// This is an example of a simple struct that can be embedded in other structs
type Customer struct {
	ID    int
	Name  string
	Email string
}

// changePrice is a method with a pointer receiver (*Orders)
// We use pointer receivers when we need to modify the struct's fields
// Using a pointer receiver is more efficient as it avoids copying the entire struct
func (O *Orders) changePrice(Price float64) {
	O.Price = Price // Direct field access through pointer (Go handles dereferencing automatically)
}

// getName is a method with a value receiver (Orders)
// We use value receivers when we only need to read data, not modify it
// This creates a copy of the struct, which is fine for read-only operations
func (O Orders) getName() string {
	return O.Name
}

// NewOrders is a constructor function for the Orders struct
// In Go, we typically use functions starting with 'New' as constructors
// This is a common pattern when you need to ensure proper initialization
func NewOrders(id int, name string, price float64) Orders {
	myOrders := Orders{
		ID:        id,
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	return myOrders
}

func main() {
	order := Orders{
		ID: 1,
		Name: "Laptop",
		Price: 999.99,
		CreatedAt: time.Now(), // Set the current time
	}
	//fmt.Println(order)

	order.ID = 2 // Update the ID
	order.Name = "Smartphone" // Update the Name
	order.Price = 499.99 // Update the Price
	order.CreatedAt = time.Now() // Update the CreatedAt time
	//fmt.Println(order)


	order2 := Orders{
		ID: 3,
		Name: "Tablet",
		Price: 299.99,
		CreatedAt: time.Now(), // Set the current time		
		}
	//fmt.Println(order2)




	order2.changePrice(199.99) // Attempt to change the price using a method
	//fmt.Println(order2) // The price will not change because the method does not modify the original

	//fmt.Println(order2.getName())


	// Create a new order using the constructor-like function
	newOrder := NewOrders(4, "Smartwatch", 199.99)
	fmt.Println(newOrder) // This will print the new order with the correct values





	//if struct is used only one time then
	language := struct{
		Name string
		Age int
	}{
		Name: "Go",
		Age: 10,
	}
	fmt.Println(language)




	newOrder1 := Orders{
		ID: 5,
		Name: "Headphones",
		Price: 199.99,
		CreatedAt: time.Now(),
		Customer: Customer{
			ID: 1,
			Name: "John Doe",
			Email: "john.doe@example.com",
		},
	}
	fmt.Println(newOrder1) // This will print the new order with the customer information
	newOrder1.Customer.Name = "rishabh"
	fmt.Println(newOrder1.Customer.Name) // This will print the updated customer name
}
