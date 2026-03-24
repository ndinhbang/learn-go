package main

import (
	"fmt"
	"time"
)

func infiniteCount(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Second * 1)
	}
}

// Using goroutines: "dog" is counted in the background.
func main() {
	go infiniteCount("dog")
	infiniteCount("cat")
}
