package main

import (
	"fmt"
	"time"
)

func main(){
	// x := 3
	// switch x{
	// 	case 1:
	// 		fmt.Println("x is 1")
	// 	case 2:
	// 		fmt.Println("x is 2")
	// 	case 3:
	// 		fmt.Println("x is 3")
	// 	default:
	// 		fmt.Println("x is something else")
	//}


	//multiple condition switch
	
	switch time.Now().Weekday(){
		case time.Saturday, time.Sunday:
			fmt.Println("It's the weekend!")
		default:
			fmt.Println("It's a weekday")		
	}
}
