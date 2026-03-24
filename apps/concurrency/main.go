package main

import (
	"fmt"
	"sync"
	"time"
)

func infiniteCount(thing string, wg *sync.WaitGroup) {
	defer wg.Done() // make sure to call Done() when the function finishes
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Second * 1)
	}
}

// Running both count functions as goroutines.
func main() {
	var wg sync.WaitGroup
	fmt.Println("Starting goroutines...")
	wg.Add(2) // Luôn gọi wg.Add(n) trước khi khởi động các goroutine, không xen kẽ giữa chúng.
	go infiniteCount("dog", &wg)
	go infiniteCount("cat", &wg)
	// fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("All goroutines finished.")
}
