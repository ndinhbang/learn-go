package main

import (
	"fmt"
	"math/rand"
	"strings"
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

func main() {
	// Giả lập các tác vụ tốn thời gian khác nhau
	task1 := func() any {
		time.Sleep(200 * time.Millisecond)
		return "Task 1 xong"
	}
	task2 := func() any {
		time.Sleep(100 * time.Millisecond)
		return "Task 2 xong"
	}
	task3 := func() any {
		time.Sleep(300 * time.Millisecond)
		return 500 // Trả về số
	}

	start := time.Now()

	// Kích hoạt các hàm async (tương đương việc tạo các Promise trong JS)
	p1 := async(task1)()
	p2 := async(task2)()
	p3 := async(task3)()

	fmt.Println("Đang chạy song song...")

	// Sử dụng PromiseAll để đợi tất cả
	allResults := await(PromiseAll(p1, p2, p3))

	elapsed := time.Since(start)

	fmt.Printf("Kết quả: %v\n", allResults)
	fmt.Printf("Tổng thời gian: %s\n", elapsed)

	// Giải thích: Tổng thời gian sẽ rơi vào khoảng 300ms (thời gian của task lâu nhất)
	// thay vì 600ms (tổng thời gian 3 task cộng lại).
}
