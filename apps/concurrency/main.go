package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
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

func main() {
	stream := make(chan bool)

	send := func() {
		fmt.Println("sender: ready to send...")
		stream <- true // (1)
		fmt.Println("sender: sent!")
	}

	receive := func() {
		fmt.Println("receiver: not ready yet...")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("receiver: ready to receive...")
		<-stream // (2)
		fmt.Println("receiver: received!")
	}

	var wg sync.WaitGroup
	wg.Go(send)
	wg.Go(receive)
	wg.Wait()
}
