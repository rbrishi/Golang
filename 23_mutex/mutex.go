package main

/*
COMPREHENSIVE GUIDE TO MUTEX IN GO

KEY CONCEPTS AND DEFINITIONS:

1. Mutex (Mutual Exclusion):
   - A synchronization primitive used to protect shared resources
   - Ensures only one goroutine can access protected data at a time
   - Prevents race conditions in concurrent programming
   - Part of the sync package

2. Race Condition:
   - Occurs when multiple goroutines access shared data concurrently
   - Can lead to unpredictable behavior and data corruption
   - Common in concurrent programs without proper synchronization
   - Example: Two goroutines trying to increment a counter simultaneously

3. Critical Section:
   - Section of code that accesses shared resources
   - Must be protected from concurrent access
   - Wrapped between mutex Lock() and Unlock() calls
   - Should be as small as possible for better performance

4. sync.WaitGroup:
   - Used to wait for a collection of goroutines to finish
   - Add(n): Adds n to the counter
   - Done(): Decrements counter by 1
   - Wait(): Blocks until counter becomes 0

5. Lock() and Unlock():
   - Lock(): Acquires exclusive access to the protected resource
   - Unlock(): Releases the lock for other goroutines
   - Must always come in pairs
   - defer is commonly used with Unlock()

6. Best Practices:
   - Always unlock mutex in a defer statement
   - Keep critical sections as small as possible
   - Avoid nested locks to prevent deadlocks
   - Use mutex as a struct field to protect that struct's fields
   - Document which fields are protected by the mutex

7. Common Patterns:
   - Protecting shared state in structs
   - Ensuring atomic operations
   - Implementing thread-safe data structures
   - Coordinating access to resources
*/

import (
	"fmt"
	"sync"
)

// Post represents a blog post with a view counter
// This struct demonstrates a common pattern of embedding a mutex
// to protect access to other fields in the struct
type Post struct{
	Views int        // The number of views, protected by mu
	mu sync.Mutex    // Mutex to protect the Views field from concurrent access
}

// inc increments the view count of a Post in a thread-safe manner
// It uses both a mutex for protecting the Views field and a WaitGroup for synchronization
// Parameters:
//   - wg: A pointer to sync.WaitGroup for coordinating multiple goroutines
func (P *Post)inc(wg *sync.WaitGroup) {
	defer func(){
		P.mu.Unlock() // Ensure the mutex is unlocked when the function returns
		wg.Done() // Decrement the WaitGroup counter when the goroutine completes
	}()
	P.mu.Lock()         // Acquire the mutex lock
	P.Views += 1
}

func main(){
	// Create a new Post instance with initial view count of 0
	myPost := Post{Views: 0}

	// Create a WaitGroup to synchronize multiple goroutines
	var wg sync.WaitGroup
	
	// Add 100 to the WaitGroup counter
	// This tells the WaitGroup to wait for 100 goroutines to complete
	wg.Add(100)

	// Launch 100 goroutines that will concurrently increment the view count
	// This demonstrates:
	// 1. How race conditions could occur without mutex
	// 2. How mutex prevents these race conditions
	// 3. How WaitGroup coordinates multiple goroutines
	for i := 0; i < 100; i++ {
		go myPost.inc(&wg)  // Each goroutine calls inc() method
	}

	// Wait for all goroutines to complete
	// This ensures we don't print the result until all increments are done
	wg.Wait()

	// Print the final view count
	// Without mutex, this could show less than 100
	// With mutex, it will always show exactly 100
	fmt.Println("Views:", myPost.Views)
}
