package main

import "fmt"


var add string = "hello world" // package level variable, can be accessed in the entire package
var add2 = "hello world 2" // package level variable, can be accessed in the entire package
//add3 := "hello world 3" // short variable declaration, can not be used at package level, only inside functions

func main() {
	//var name string  = "rishabh"
	var name  = "rishabh"
	// var age int = 20
	var age = 20


	//short variable declaration 	
	name1 := "rishabh"
	age1 := 20 

	//case 3
	var name3 string
	name3 = "bajpai"

	fmt.Println(add, add2, name, age, name1, age1, name3)
}