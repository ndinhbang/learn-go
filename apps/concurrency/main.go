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

	for word := range in {
		fmt.Printf("%s ", word)
	}
}
