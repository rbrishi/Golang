package main

import "fmt"

func main(){
	const pi = 3.14
	const name = "Rishabh Bajpai"
	const age int = 20 // const age = 20 // this will also work, but type is not specified
	//age = 21 // this will not work, as const cannot be reassigned

	//constangt grouping
	// constants can be grouped together using parentheses
	// this is useful for better organization and readability of code
	// constants can be of different types, but they must be of the same type in a
	const (
		port = 8080
		host = "localhost"
		dbName = "mydb"
	)


	fmt.Println(pi, name, age, port, host, dbName)
}