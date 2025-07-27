package main

import (
	"fmt"
	"slices"
)

func main(){
	//uninitialized slice is nil
	var slice []int
	fmt.Println(slice)
	fmt.Println(slice == nil, len(slice), cap(slice)) // true 0 0

	//var slice2 []int = make([]int, 0) // slice with length 0 and capacity 0
	//fmt.Println(slice2, slice2 == nil, len(slice2), cap(slice2)) // [] true 0 0
	var slice3 = make([]int, 2, 10) // slice with length 2 and capacity 10
	fmt.Println(slice3, slice3 == nil, len(slice3), cap(slice3)) // [0 0] false 2 10


	slice3 = append(slice3, 1, 2, 3, 4, 5) // appending elements
	fmt.Println(slice3, slice3 == nil, len(slice3), cap(slice3)) // [0 0 1 2 3 4 5] false 7 10



	//3rd way to create a slice
	slice4 := []int{}
	fmt.Println(slice4)
	slice4 = append(slice4, 1,2,3)
	fmt.Println(slice4, len(slice4), cap(slice4))
	slice4[0] = 5
	fmt.Println((slice4))



	//copy slice
	num1 := []int{11,12,13}
	num2 := make([]int, len(num1))
	copy(num2, num1) // destination, source
	fmt.Println(num1, num2)



	//slice operator
	num3 := []int{1,2,3,4}
	fmt.Println(num3[0:3]) //[1 2 3]
	fmt.Println(num3[:2]) // [1 2]
	fmt.Println(num3[1:]) // [2 3 4]

	fmt.Println(slices.Equal(num1, num3)) // false
	//to access slices package use "slices."



	//2d slice
	num4 := [][]int{{0,1,2}, {3,4,5}}
	fmt.Println(num4) // [[0 1 2] [3 4 5]]

}
