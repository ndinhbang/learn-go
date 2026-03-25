package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// only send to the channel
func submit(str string, stream chan<- string) {
	words := strings.Split(str, ",")
	for _, word := range words {
		stream <- word
	}
	close(stream)
}

// only receive from the channel
func receive(stream <-chan string) {
	for word := range stream {
		fmt.Printf("%s ", word)
	}
	fmt.Println()
}

// speaks a phrase word by word with some pauses
func say(done chan struct{}, id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
	done <- struct{}{} // send a signal to the channel to indicate that the phrase is done
}

func work(done chan<- struct{}, out chan<- int) {
	for i := 1; i <= 5; i++ {
		out <- i
	}
	// close(out)
	done <- struct{}{}
	// close(done)
}

// async converts a regular function to an asynchronous one.
// An asynchronous function returns a result channel when called.
func async(fn func() any) func() chan any {
	return func() chan any {
		out := make(chan any, 1)
		go func() {
			out <- fn()
		}()
		return out
	}
}

// await waits for the result of an asynchronous function
// on the given channel.
func await(in <-chan any) any {
	return <-in
}

// PromiseAll nhận vào một slice các channels và trả về một channel chứa slice kết quả
func PromiseAll(channels ...<-chan any) <-chan any {
	out := make(chan any, 1)

	go func() {
		results := make([]any, len(channels))

		// Sử dụng một cơ chế để đợi tất cả các goroutine hoàn thành
		// Ở đây ta có thể đọc trực tiếp từ từng channel theo thứ tự
		for i, ch := range channels {
			results[i] = <-ch // Đợi từng channel trả về kết quả
		}

		out <- results // Gửi mảng kết quả cuối cùng vào channel đầu ra
		close(out)
	}()

	return out
}

func workerTask(id int, p string) {
	fmt.Printf("[Worker %d] Đang xử lý: %s\n", id, p)
	time.Sleep(1 * time.Second) // Giả lập công việc nặng
}

func main() {
	phrases := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	var wg sync.WaitGroup
	maxConcurrency := 3
	sema := make(chan struct{}, maxConcurrency) // Dùng struct{} để tiết kiệm bộ nhớ

	for i, p := range phrases {
		wg.Add(1)
		sema <- struct{}{} // "Lấy chỗ" - Nếu đủ 3 chỗ rồi, nó sẽ đợi ở đây

		go func(id int, phrase string) {
			defer wg.Done()           //
			defer func() { <-sema }() // "Trả chỗ" khi làm xong

			workerTask(id, phrase)
		}(i, p)
	}

	wg.Wait()
	fmt.Println("All workers done")
}
