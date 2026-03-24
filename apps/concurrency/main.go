package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

// Create a function countDigitsInWords that takes an input string,
// splits it into words, and counts the digits in each word using countDigits.
// Be sure to do the counting for each word in a separate goroutine.

// countDigits returns the number of digits in a string.
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// counter stores the number of digits in each word.
// The key is the word, and the value is the number of digits.
type counter map[string]int

func countDigitsInWords(phrase string) counter {
	countChan := make(chan int)
	words := strings.FieldsSeq(phrase) // Split the phrase into words using FieldsSeq to handle multiple spaces.

	for word := range words {
		go func() {
			count := countDigits(word)
			countChan <- count
		}()
	}

	stats := counter{}

	for word := range words {
		count := <-countChan
		stats[word] = count
	}

	return stats
}

// asStats converts statistics from sync.Map to a regular map.
func asStats(m *sync.Map) counter {
	stats := counter{}
	m.Range(func(word, count any) bool {
		stats[word.(string)] = count.(int)
		return true
	})
	return stats
}

// printStats prints the number of digits in words.
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

// Running both count functions as goroutines.
func main() {
	str := "one,two,,four"

	in := make(chan string)
	go func() { // (1)
		words := strings.Split(str, ",")
		for _, word := range words {
			in <- word
		}
	}()

	for { // (2)
		word := <-in
		if word != "" {
			fmt.Printf("%s ", word)
		}
	}
	// one two four
}
