package main

import (
	"fmt"
	"log"

	"github.com/ndinhbang/learn-go/pkg/greetings"
)

func main() {
	// Get a greeting message and print it.
	message, err := greetings.Hello("Bang")

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
