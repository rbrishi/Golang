package main

import "fmt"

//pass by value
// func changeNum(num int){
// 	num = 10
// 	fmt.Println("Inside changeNum:", num)
// }

// pass by reference
func changeNum(num *int) {
	*num = 10
	fmt.Println("Inside changeNum:", *num)
}

func main() {
	num := 1
	fmt.Println("Before changeNum:", num)
	changeNum(&num)
	fmt.Println("After changeNum:", num)
}

// * is used to dereference the pointer
// & is used to get the address of the variable	