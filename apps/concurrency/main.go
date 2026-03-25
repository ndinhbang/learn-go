package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
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

// speaks a phrase word by word with some pauses
func say(done chan struct{}, id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
	done <- struct{}{} // send a signal to the channel to indicate that the phrase is done
}

func work(done chan struct{}, out chan int) {
	for i := 1; i <= 5; i++ {
		out <- i
	}
	done <- struct{}{}
}

func main() {
	out := make(chan int)
	done := make(chan struct{})

	go work(done, out) // (1)

	<-done // (2)

	for n := range out { // (3)
		fmt.Println(n)
	}
}
