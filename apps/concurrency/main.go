package main

import (
	"fmt"
	"sync"
	"time"
)

func infiniteCount(thing string, wg *sync.WaitGroup) {
	defer wg.Done() // Indicates that this goroutine is done when the function returns.
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
	go infiniteCount("dog", &wg)
	go infiniteCount("cat", &wg)
	// fmt.Println("Waiting for goroutines to finish...")
	wg.Wait() // Wait for all goroutines to finish before proceeding to the next line of code.
	fmt.Println("All goroutines finished.")
}
