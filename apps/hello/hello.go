package main

import (
	"fmt"

	"github.com/ndinhbang/learn-go/pkg/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Bang")
	fmt.Println(message)
}
