package main

import (
	"fmt"
	"maps"
)

func main() {
	// Method 1: Creating a map using make
	// Syntax: make(map[KeyType]ValueType)
	// This is preferred when you don't know the initial values
	// The size is optional and the map will grow dynamically
	m := make(map[string]int)

	// Adding key-value pairs to the map
	m["foo"] = 1
	m["bar"] = 2

	// Accessing map values
	fmt.Println("Value for key 'foo':", m["foo"]) // Output: Value for key 'foo': 1
	fmt.Println("Entire map:", m)                 // Output: Entire map: map[bar:2 foo:1]

	// Safe way to check if a key exists in the map
	// The second return value (ok) is a boolean indicating if the key was found
	value, ok := m["bar"]
	if !ok {
		fmt.Println("Key 'bar' not found in map")
	} else {
		fmt.Println("Value for key 'bar':", value)
	}

	// Default values when key doesn't exist:
	// - string: empty string ("")
	// - bool: false
	// - numeric types: 0
	// - pointer, slice, map, interface: nil
	// - struct: empty struct

	// Map operations
	fmt.Printf("Map length: %d\n", len(m)) // Output: Map length: 2

	// Deleting a key from the map
	delete(m, "foo")
	fmt.Println("After deleting 'foo':", m) // Output: After deleting 'foo': map[bar:2]

	// Clearing all elements from the map (Go 1.21+)
	clear(m)
	fmt.Println("After clearing:", m) // Output: After clearing: map[]

	// Method 2: Creating a map using map literal
	// This is preferred when you know the initial values
	scores := map[string]int{
		"Alice": 95,
		"Bob":   87,
		"Carol": 92,
	}

	// Iterating over a map
	fmt.Println("\nStudent Scores:")
	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}

	// Iterating over just keys
	fmt.Println("\nStudent Names:")
	for name := range scores {
		fmt.Printf("- %s\n", name)
	}

	// Note: Map iteration order is not guaranteed in Go
	// The order may be different each time you run the program

	// Map of maps (nested maps)
	students := map[string]map[string]int{
		"Alice": {"Math": 95, "Science": 98},
		"Bob":   {"Math": 87, "Science": 89},
	}

	// Safely accessing nested maps
	if mathScore, ok := students["Alice"]["Math"]; ok {
		fmt.Printf("\nAlice's Math score: %d\n", mathScore)
	}


	//check both maps are equal or not
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 2, "a": 1}

	fmt.Println(maps.Equal(m1, m2)) // Output: true
}
