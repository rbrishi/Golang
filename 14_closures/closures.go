package main

import "fmt"

// counter is a function that demonstrates the concept of closures in Go
// A closure is a function value that references variables from outside its body
// It returns a function that has access to the count variable even after counter() completes
func counter() func() int {
	// count is initialized in the outer function
	// This variable will be "closed over" by the inner function
	count := 0

	// Return an anonymous function that has access to count
	// This creates a closure where the returned function "closes over" the count variable
	return func() int {
		count += 1 // Modifies the count variable from the outer scope
		return count
	}
}

func main() {
	// Create a new counter
	// increment is assigned the inner function returned by counter()
	// It maintains its own count variable separate from any other counters
	increment := counter()

	// Each time we call increment(), it increases its own count by 1
	fmt.Println(increment()) // Output: 1
}