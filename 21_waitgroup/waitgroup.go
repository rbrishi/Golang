package main

// The sync package provides synchronization primitives like WaitGroup
import (
	"fmt"
	"sync"
)

// WaitGroup is used to wait for a collection of goroutines to finish executing
// It works like a counter:
// - wg.Add(n) increments the counter by n
// - wg.Done() decrements the counter by 1
// - wg.Wait() blocks until the counter becomes 0

// task simulates a concurrent operation
// Parameters:
//   - id: identifier for the task
//   - wg: pointer to WaitGroup for synchronization
func task(id int, wg *sync.WaitGroup){
	// defer ensures wg.Done() is called when the function returns
	// This decrements the WaitGroup counter
	defer wg.Done()
	fmt.Println("Task", id, "started")
}

func main(){
	// Create a WaitGroup instance to synchronize goroutines
	var wg sync.WaitGroup 

	// Launch 11 goroutines (0 to 10)
	for i := 0; i <= 10; i++{
		// Increment WaitGroup counter before launching each goroutine
		wg.Add(1)

		// Launch task as a goroutine
		// 'go' keyword creates a new goroutine (lightweight thread)
		// Each task runs concurrently with other goroutines
		go task(i, &wg)
	}

	// Wait blocks here until all goroutines complete
	// (until WaitGroup counter becomes 0)
	wg.Wait()
}
