package main

import "fmt"

// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//generic function 
func printSlice[T any](items []T){
	for _, item := range items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}
//func printSlice[T interface{]}](items []T)
// func printSlice[T int | string](items []T)
//func printSlice[T comparable](items []T)


//generic struct
type Stack[T any] struct {
	elements []T 
}

func main(){
	printSlice([]int{1, 2, 3})
	printSlice([]string{"a", "b", "c"})

	myStack := Stack[int]{elements: []int{1, 2, 3}}
	fmt.Println(myStack)
}
