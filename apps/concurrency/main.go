package main

import (
	"fmt"
	"sync"
	"time"
)

func infiniteCount(thing string) {

	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Second * 1)
	}
}

// Running both count functions as goroutines.
func main() {
	var wg sync.WaitGroup
	fmt.Println("Starting goroutines...")
	// The WaitGroup.Go method (Go 1.25+) is used to start a goroutine and automatically manage the WaitGroup counter.
	// It increments the counter when the goroutine starts and decrements it when the goroutine finishes.
	// This eliminates the need for manual calls to wg.Add(1) and defer wg.Done() in each goroutine.
	wg.Go(func() {
		infiniteCount("dog")
	})
	wg.Go(func() {
		infiniteCount("cat")
	})
	// fmt.Println("Waiting for goroutines to finish...")
	wg.Wait() // Wait for all goroutines to finish before proceeding to the next line of code.
	fmt.Println("All goroutines finished.")
}
