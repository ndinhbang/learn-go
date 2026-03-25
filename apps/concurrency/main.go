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

func work(done chan<- struct{}, out chan<- int) {
	for i := 1; i <= 5; i++ {
		out <- i
	}
	// close(out)
	done <- struct{}{}
	// close(done)
}

// async converts a regular function to an asynchronous one.
// An asynchronous function returns a result channel when called.
func async(fn func() any) func() <-chan any {
	return func() <-chan any {
		out := make(chan any, 1)
		go func() {
			out <- fn()
		}()
		return out
	}
}

// await waits for the result of an asynchronous function
// on the given channel.
func await(in <-chan any) any {
	return <-in
}
func main() {
	fn := func() any {
		time.Sleep(100 * time.Millisecond)
		return "okay"
	}

	slowpoke := async(fn) // create an asynchronous function

	start := time.Now()
	slowpoke()                  // does not block
	slowpoke()                  // does not block
	slowpoke()                  // does not block
	result := await(slowpoke()) // blocks until the result is ready

	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println("took", elapsed)
	// okay
	// took 100ms

	// total execution time is 100ms, not 400ms
}
