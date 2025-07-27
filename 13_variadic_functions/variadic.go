package main

import "fmt"

// Demonstration of variadic functions in Go
// Variadic functions can accept a variable number of arguments

// sum is a variadic function that takes a variable number of int arguments
// It sums all the integers passed to it and returns the total
// Parameters:
//   - num ...int: Variable number of integers
// Returns:
//   - int: Sum of all provided integers
func sum(num ...int) int {
	total := 0
	for _, n := range num {
		total += n
	}
	return total
}

// addition demonstrates type assertion in Go using variadic parameters
// It accepts arguments of any type (interface{}) and identifies their specific types
// Parameters:
//   - num ...interface{}: Variable number of arguments of any type
// The function uses if-else for type assertion instead of type switch (shown in comments)
func addition(num ...interface{}) {
	for _, n := range num {
		// Alternative implementation using type switch (commented out)
		// switch v := n.(type) {
		// case int:
		// 	fmt.Println("Integer:", v)
		// case string:
		// 	fmt.Println("String:", v)
		// default:
		// 	fmt.Println("Unknown type")
		// }
		
		// Type assertion using if-else
		// n.(type) attempts to assert the type of n
		// ok will be true if the type assertion is successful
		if v, ok := n.(int); ok {
			fmt.Println("Integer:", v)
		} else if v, ok := n.(string); ok {
			fmt.Println("String:", v)
		} else {
			fmt.Println("Unknown type")
		}
	}
}

func main(){
	// Example 1: Using variadic function with direct arguments
	fmt.Println(sum(1, 2, 3, 4, 5))

	// Example 2: Using addition function with multiple types
	// Demonstrates how interface{} can accept any type
	addition(1, "hello", 3.14, "world", 42, 100, "Go", 3.14159)

	// Example 3: Using variadic function with a slice
	// The ... operator unpacks the slice into individual arguments
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(nums...)) // passing a slice to a variadic function
}
