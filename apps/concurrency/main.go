package main

import (
	"fmt"
	"time"
)

// Running both count functions as goroutines.
func main() {
	messages := make(chan string)

	go func() {
		fmt.Println("B: Sending message...")
		messages <- "ping"              // (1) After sending the message to the channel, goroutine B gets blocked
		fmt.Println("B: Message sent!") // (2) This line will not be executed until goroutine A receives the message from the channel
	}()

	fmt.Println("A: Doing some work...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("A: Ready to receive a message...")

	<-messages //  (3) Goroutine A receives the message from the channel, unblocking goroutine B and allowing it to continue its execution. After receiving the message, goroutine A can proceed with its work.

	fmt.Println("A: Messege received!")
	time.Sleep(100 * time.Millisecond)

}
