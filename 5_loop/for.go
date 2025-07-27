package main

import "fmt"

func main(){
	//while loop
	i := 0
	for i < 5{
		fmt.Println(i)
		i++
	}

	//infinite loop
	// for{
	// 	fmt.Println("infinite loop")
	// }


	//classic for loop
	for i := 0; i < 5; i++{
		fmt.Println(i)
	}



	//range -> go version 1.22
	for  i := range 3{
		fmt.Println(i)
	} 
}
