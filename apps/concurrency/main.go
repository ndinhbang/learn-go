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
	wg.Add(2) // Add the number of goroutines we are going to wait for (2 in this case).
	go func() {
		defer wg.Done() // Indicates that this goroutine is done when the function returns.
		infiniteCount("dog")
	}()

	go func() {
		defer wg.Done() // Indicates that this goroutine is done when the function returns.
		infiniteCount("cat")
	}()
	// fmt.Println("Waiting for goroutines to finish...")
	wg.Wait() // Wait for all goroutines to finish before proceeding to the next line of code.
	fmt.Println("All goroutines finished.")
}
