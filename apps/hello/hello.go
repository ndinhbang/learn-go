package main

import (
	"fmt"
	"sync"
)

func say(s string, wg *sync.WaitGroup) {
	if wg != nil {
		// Hãy đợi cho đến khi hàm này thực hiện xong tất cả mọi thứ, rồi mới thực hiện lệnh này ngay trước khi thoát ra
		defer wg.Done() // Báo cáo đã xong việc khi hàm kết thúc (giảm đếm 1)
	}
	for range 5 {
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup

	// // Get a greeting message and print it.
	// message, err := greetings.Hello("Bang")

	// // If an error was returned, print it to the console and
	// // exit the program.
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // If no error was returned, print the returned message
	// // to the console.
	// fmt.Println(message)
	wg.Add(1)               // Thông báo: Có 1 Goroutine chuẩn bị chạy
	go say("Thế giới", &wg) // Chạy ở nền (background), Chỉ Goroutine này mới cần WaitGroup
	say("Chào", nil)        // Chạy ở luồng chính (main), truyền nil để hàm không gọi Done()

	wg.Wait() // Dừng tại đây, đợi cho đến khi bộ đếm về 0
}
