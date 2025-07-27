package main

import "fmt"

func main(){
	age := 11
	
	if age >= 18{
		fmt.Println("You are an adult")
	}else{
		fmt.Println("You are a minor")
	}

	//we can also declare variables in the if statement
	if name := "rishabh"; name  == "rishabh" {
		fmt.Println("Hello Rishabh")
	} else if name == "john" {
		fmt.Println("Hello John")
	}


	//go does not have a ternary operator like other languages
	// so we use if else statements instead
}
