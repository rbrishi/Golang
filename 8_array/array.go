package main

import (
	"fmt"
)

func main(){
	var arr [5]int // Declare an array of 5 integers
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr)) // Print the length and capacity of the array

	arr[0] = 10 
	arr[1] = 20 
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50 // Assign values to the array
	fmt.Println(arr) 
	fmt.Println(len(arr)) 
	fmt.Println(cap(arr))
	//arr[5] = 60 // This will cause a runtime error: index out of range


	var strArr [3]string // Declare an array of 3 strings
	strArr[0] = "Hello"
	fmt.Println(strArr) // Print the string array

	var nums [5]int = [5]int{1, 2, 3, 4, 5} // Declare and initialize an array
	fmt.Println(nums) // Print the initialized array


	nums2 := [5]int{6, 7, 8, 9, 10} // Declare and initialize another array
	fmt.Println(nums2) // Print the second initialized array


	//2d array
	var matrix [2][2]int // Declare a 2D array with 2 rows and 2 columns
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[1][0] = 3
	matrix[1][1] = 4
	fmt.Println(matrix) // Print the 2D array


	matrix2 := [2][2]int{
		{5, 6},
		{7, 8},
	} // Declare and initialize another 2D array
	fmt.Println(matrix2) // Print the second 2D array


	var matrix3 [2][2]int = [2][2]int{ // Declare and initialize a 2D array with different values
		{9, 10},
		{11, 12},
	}
	fmt.Println(matrix3) // Print the third 2D array

	fmt.Println(len(matrix3)) // Print the length of the 2D array
	fmt.Println(cap(matrix3)) // Print the capacity of the 2D array	
}