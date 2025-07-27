package main

import (
	"fmt"
)

// The range keyword in Go is used for iterating over various data structures:
// - Arrays and Slices: returns index and value
// - Maps: returns key and value
// - Strings: returns index and rune (Unicode code point)
// - Channels: returns values received from the channel
func main()  {
	// Example 1: Using range with a slice
	// Creating a slice of integers
	nums := []int{10,20,30,40,50}

	// Traditional for loop (commented out for reference)
	// for i := 0; i < len(nums); i++{
	// 	fmt.Println(nums[i]) // Output: 10 20 30 40 50
	// }


	// Example 2: Range with slices returns both index and value
	for index, value := range nums {
		fmt.Println("Index:", index, "Value:", value)
	}

	// Example 3: Using blank identifier (_) to ignore the index
	// Useful when you only need the values
	sum := 0
	for _, num := range nums{
		sum += num
	}
	fmt.Println("Sum of nums:", sum) // Output: Sum of nums: 150



	// Example 4: Using range with maps
	// Creating a map with string keys and string values
	m := map[string]string{
		"one": "10",
		"two": "20",
	}

	// Range over map returns both key and value in each iteration
	// Note: Map iteration order is not guaranteed in Go
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	// If you use only one variable in range with maps:
	// - It will be assigned the key (not the value)
	// - This is different from arrays/slices where single variable gets the value
	for i := range m {
		fmt.Println("Key:", i, "Value:", m[i])
	}

	// Example 5: Using range with strings
	// When ranging over a string, range provides:
	// - index: byte index of each rune
	// - char: rune value (Unicode code point)
	str := "hello"
	for index, char := range str {
		// Using Printf with %c format specifier to print character
		fmt.Printf("Index: %d, Character: %c\n", index, char)
		
		// Two ways to print the character:
		// 1. string(char): converts the rune to a string
		// 2. Just 'char' would print the Unicode value (e.g., 104 for 'h')
		fmt.Println("Index:", index, "Character:", string(char))
	}
}
