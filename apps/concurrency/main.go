package main

import (
	"fmt"
	"strings"
)

// Running both count functions as goroutines.
func main() {
	str := "one,two,,four"

	in := make(chan string)
	go func() { // (1)
		words := strings.Split(str, ",")
		for _, word := range words {
			in <- word
		}
		// close the channel to signal that there are no more words
		// only one goroutine should close the channel
		// if multiple goroutines close the channel, it will panic
		// so we need to make sure that only one goroutine closes the channel
		// and the other goroutines should not close the channel
		close(in)
	}()

	for { // (2)
		// receive from the channel
		word, ok := <-in
		// if the channel is closed, break the loop
		// if the channel is closed, the ok will be false
		// if the channel is not closed, the ok will be true
		if !ok {
			break
		}
		fmt.Printf("%s ", word)
	}
}
