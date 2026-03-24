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

// Running both count functions as goroutines.
func main() {
	go infiniteCount("dog")
	go infiniteCount("cat")

}
