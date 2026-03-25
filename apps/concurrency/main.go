package main

import (
	"fmt"
	"strings"
)

// only send to the channel
func submit(str string, stream chan<- string) {
	words := strings.Split(str, ",")
	for _, word := range words {
		stream <- word
	}
	close(stream)
}

// only receive from the channel
func receive(stream <-chan string) {
	for word := range stream {
		fmt.Printf("%s ", word)
	}
	fmt.Println()
}

// Running both count functions as goroutines.
func main() {
	str := "one,two,,four"

	stream := make(chan string)
	go submit(str, stream)
	receive(stream)
}
