package main

import (
	"fmt"
)

// Running both count functions as goroutines.
func main() {
	msgChannel := make(chan string)

	go func() {
		msgChannel <- "Hello from the goroutine!"
	}()
	fmt.Println("Waiting for message from goroutine...")
	msg := <-msgChannel // blocks until a message is received from the channel
	fmt.Println("Received message from goroutine:")
	fmt.Println(msg)

}
