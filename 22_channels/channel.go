package main

/*
COMPREHENSIVE GUIDE TO GO CHANNELS

KEY CONCEPTS AND DEFINITIONS:

1. Goroutine:
   - A lightweight thread of execution managed by the Go runtime
   - Created using the 'go' keyword
   - Runs concurrently with other goroutines
   - Much cheaper than OS threads

2. Channel:
   - A typed conduit for communication between goroutines
   - Created using 'make(chan Type)'
   - Can be buffered or unbuffered
   - Provides synchronized data transfer

3. Buffer:
   - A temporary storage area for channel messages
   - Specified during channel creation: make(chan Type, bufferSize)
   - Allows asynchronous communication up to buffer capacity
   - Helps manage back-pressure in concurrent systems

4. Select Statement:
   - Like a switch for channel operations
   - Allows handling multiple channel communications
   - Can include default case for non-blocking operations
   - Randomly selects when multiple cases are ready

5. Keywords and Operators:
   - make(): Built-in function to create channels
   - <- : Channel operator (send/receive)
   - close(): Function to close a channel
   - range: Keyword to iterate over channel values
   - defer: Delays execution until surrounding function returns

6. Channel States:
   - Open: Normal operation state
   - Closed: No more values can be sent
   - Nil: Zero value of channel type, blocks forever

7. Synchronization Patterns:
   - WaitGroup: For managing multiple goroutines
   - Done Channel: For signaling completion
   - Worker Pool: For parallel task processing
   - Fan-out/Fan-in: For distributing/collecting work

8. Best Practices:
   - Channel owner should close the channel
   - Don't close receive-only channels
   - Don't close already closed channels
   - Check closed status when receiving
   - Use buffered channels judiciously
   - Always handle potential deadlocks

9. Common Channel Operations:
   - Send: ch <- value
   - Receive: value := <-ch
   - Close: close(ch)
   - Check if closed: value, ok := <-ch

This file demonstrates various aspects of channel usage:

1. Basic Channel Operations:
   - Creating channels using make(chan Type)
   - Sending data to channels using ch <- value
   - Receiving data from channels using <-ch
   - Blocking nature of channels

2. Channel Types:
   - Unbuffered channels (synchronous communication)
   - Buffered channels (asynchronous communication with capacity)
   - Receive-only channels (<-chan Type)
   - Send-only channels (chan<- Type)

3. Channel Patterns:
   - Goroutine synchronization
   - Worker pools
   - Signal channels (done channel pattern)
   - Range over channels
   - Select statement for multiple channels

4. Best Practices:
   - Proper channel closure
   - Error handling
   - Avoiding deadlocks
   - Channel direction specification
*/

import (
	"fmt"
	"time"
	//"time"
	//"math/rand"
)

// processNum demonstrates basic channel usage
// It receives an integer from a channel and prints it
// This shows the simplest form of channel communication
func processNum(numChan chan int){
		fmt.Println(<-numChan) // Receive a number from the channel
}


// sum demonstrates using channels for computation results
// It takes two integers, adds them, and sends the result through a channel
// This pattern is useful when you want to perform concurrent computations
func sum(result chan int, num1 ,num2 int){
	result <- num1 + num2
}

// task demonstrates the "done" channel pattern for goroutine synchronization
// It uses a boolean channel to signal when the task is complete
// The defer statement ensures the done signal is sent even if the function panics
func task(done chan bool){
	defer func(){done <- true}()
	fmt.Println("Task is running")
}



// emailSender demonstrates several important channel concepts:
// 1. Using range to iterate over a channel until it's closed
// 2. Combining done channel pattern with a work channel
// 3. Simulating real-world async operations (email sending)
// 4. Proper cleanup using defer
func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("Sending email:", email)
		time.Sleep(time.Second) // Simulate time taken to send email
	}
}







func main(){
	// SECTION 1: BASIC CHANNEL CONCEPTS
	//
	// Channels are Go's built-in mechanism for goroutine communication
	// Key characteristics:
	// - Channels provide synchronized communication
	// - They are blocking by default (unbuffered channels)
	// - Send operation (<-) blocks until there's a receiver
	// - Receive operation blocks until there's a sender
	//
	// Example of creating a channel:
	//messageChan := make(chan string)
	
	//messageChan <- "ping" // Send a message to the channel

	//msg := <-messageChan // Receive a message from the channel
	//fmt.Println("Received:", msg)



	numChan := make(chan int)
	go processNum(numChan)
	// numChan <- 42
	// time.Sleep(time.Second)

	// for{
	// 	numChan <- rand.Intn(100) // Send a random number to the channel
	// }


	result := make(chan int)
	go sum(result, 10, 20)
	ans := <-result
	fmt.Println(ans)



	done := make(chan bool) //unbuffered channel
	// A channel is a way to communicate between goroutines
	// It can be used to send and receive messages between goroutines
	go task(done)
	<-done // bloacking until the task is done 
	// Alternatively, you can use a select statement to wait for multiple channels
	fmt.Println("Task is done")


	//Note: for single goroutine, you can use a channel
	//but for multiple goroutines, you can use a WaitGroup
	//WaitGroup is used to wait for a collection of goroutines to finish



	// Note: If you want to send a message to a channel without blocking, you can use a buffered channel
	// Buffered channels allow you to send messages without blocking until the buffer is full
	// For example, you can create a buffered channel with a capacity of 10:
	bufferedChan := make(chan string, 10)
	// You can then send messages to the buffered channel without blocking
	bufferedChan <- "message 1"
	bufferedChan <- "message 2"
	bufferedChan <- "message 3"
	fmt.Println(<-bufferedChan) // Receive a message from the buffered channel
	fmt.Println(<-bufferedChan) // Receive another message from the buffered channel
	// When the buffer is full, sending will block until there is space available
	close(bufferedChan)




	emailChan := make(chan string, 50)
	done1 := make(chan bool)
	go emailSender(emailChan, done1)
	for i:= 1; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@example.com", i) //Sprintf used for formatting strings 
	}
	fmt.Println("All emails sent")
	close(emailChan) //it is important to close the channel to signal that no more messages will be sent
	<-done1




	// SECTION 3: SELECT STATEMENT
	//
	// Select provides a way to handle multiple channel operations:
	// - Similar to switch but for channel operations
	// - Can handle both send and receive operations
	// - If multiple cases are ready, one is chosen randomly
	// - Can include a default case for non-blocking operations
	// - Useful for timeouts and channel orchestration

	chan1 := make(chan int)
	chan2 := make(chan string)
	go func() {chan1 <- 10}()
	go func() {chan2 <- "hello"}()

	//to receive data from multiple channels, we can use a combination of for and select statement
	for i := 0; i < 2; i++ {
		select {
		case num := <-chan1:
			fmt.Println("Received from chan1:", num)
		case str := <-chan2:
			fmt.Println("Received from chan2:", str)
		}	
	}
}



// CHANNEL DIRECTION SPECIFICATIONS
//
// In Go, we can specify channel direction to make our APIs more precise:
// 1. chan T     - bidirectional channel (can send and receive)
// 2. chan<- T   - send-only channel
// 3. <-chan T   - receive-only channel
//
// Example of a receive-only channel:
// func receiveOnlyChannel(ch <-chan int) {
// 	// This function can only receive from the channel
// 	// It cannot send to the channel
// 	num := <-ch
// 	fmt.Println("Received from receiveOnlyChannel:", num)
// }
//
// Benefits of channel direction specification:
// - Better documentation
// - Compile-time safety
// - Clearer API contracts
// - Prevention of accidental misuse
