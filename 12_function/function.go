// Package main is the entry point for this Go program demonstrating various function concepts
package main

// fmt package is imported for printing and formatting
import "fmt"

// add is a simple function that demonstrates:
// - Basic function declaration syntax
// - Multiple parameters of the same type (int)
// - Single return value
// - Parameters a and b are added and result is returned
func add(a, b int) int {
	return a + b
}

// getLanguage demonstrates multiple return values in Go
// Returns:
// - Three strings representing programming languages
// - A boolean indicating popularity
// This is an example of named return values and multiple return values
func getLanguage() (string, string, string, bool) {
	return "Go", "Java", "JavaScript", true		
}

// Function Concepts in Go:
// - Functions are first-class citizens in Go
// - They can be:
//   1. Passed as arguments to other functions
//   2. Returned from other functions
//   3. Assigned to variables
//   4. Stored in data structures
// processIt demonstrates how functions can be passed as parameters
// Parameters:
// - fn: a function that takes an int and returns an int
// This showcases Go's support for higher-order functions
func processIt(fn func(a int) int){
	result := fn(10)
	fmt.Println("Processed Result:", result)
}

// returnFunction demonstrates returning a function from another function
// Returns:
// - A function that takes an int parameter and returns an int
// This is an example of a closure - a function that returns another function
func returnFunction() func(a int) int{
	return func(b int) int {
		return b * 11
	}
}

// main is the entry point of the program
// It demonstrates various function concepts including:
// - Basic function calls
// - Multiple return value handling
// - Anonymous functions
// - Function variable assignment
// - Higher-order function usage
func main()  {
	// Demonstrating basic function call with add()
	result := add(10, 20)
	fmt.Println("Result:", result)

	// Demonstrating multiple return values in different ways
	// 1. Direct printing of all return values
	fmt.Println(getLanguage())
	// 2. Assigning return values to variables
	lang1, lang2, lang3, isPopular := getLanguage()
	fmt.Println("Languages:", lang1, lang2, lang3)
	fmt.Println("Is Popular:", isPopular)	

	// Demonstrating anonymous function
	// Creating a function literal and assigning it to variable fn
	fn := func(a int) int {
		return a * 2
	}
	// Passing function as parameter to processIt
	processIt(fn)

	// Demonstrating function returned from another function
	fn2 := returnFunction()
	// Using the returned function
	fmt.Println("Returned Function Result:", fn2(5)) // Output: 55

	// Alternative way to call returned function directly
	fmt.Println(returnFunction()(11))
	// Printing function value (address) - Note: This shows the memory address of the function
	fmt.Printf("%T\n", returnFunction()) // prints the type of the returned function

	// If you want to print the result of calling the returned function:
	fmt.Printf("%v\n", returnFunction()(10)) // prints the result of the function call

}
