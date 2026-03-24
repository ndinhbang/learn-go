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

// The program counts "dog" forever and never gets to "cat".
func main() {
	go infiniteCount("dog")
	infiniteCount("cat")
}
