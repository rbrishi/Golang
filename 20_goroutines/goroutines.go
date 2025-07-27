package main

import (
	"fmt"
	"time"
)

func task(id int){
	fmt.Println("Task", id, "started")
}
func main(){
	for i := 0; i <= 10; i++{
		go task(i) // output will be in random order


		go func(i int){
			fmt.Println(i)
		}(i) // Using a closure to capture the loop variable
	}
	time.Sleep(1 * time.Second) // Wait for goroutines to finish , but not a good practice

	// Using a channel or sync.WaitGroup is a better way to wait for goroutines
}