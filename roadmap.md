# Go Language - Roadmap Học Tập Chi Tiết

> Roadmap này được thiết kế cho người mới bắt đầu, giúp bạn nắm vững các khái niệm cơ bản và quan trọng trong Go.

## 📋 Mục Lục
1. [Chuẩn Bị](#chuẩn-bị)
2. [Phần 1: Cơ Bản](#phần-1-cơ-bản)
3. [Phần 2: Cấu Trúc Dữ Liệu](#phần-2-cấu-trúc-dữ-liệu)
4. [Phần 3: Xử Lý Lỗi & Defer](#phần-3-xử-lý-lỗi--defer)
5. [Phần 4: Con Con Con Con Con Con Trỏ & Memory](#phần-4-con-trỏ--memory)
6. [Phần 5: Interface & Polymorphism](#phần-5-interface--polymorphism)
7. [Phần 6: Concurrency](#phần-6-concurrency)
8. [Phần 7: Standard Library](#phần-7-standard-library)
9. [Phần 8: Testing](#phần-8-testing)
10. [Phần 9: Thực Hành Dự Án](#phần-9-thực-hành-dự-án)

---

## 🚀 Chuẩn Bị

### 1. Cài Đặt Go
- [ ] Tải và cài đặt Go từ [golang.org](https://golang.org)
- [ ] Kiểm tra cài đặt: `go version`
- [ ] Thiết lập `$GOPATH` (nếu cần)
- [ ] Hiểu cấu trúc workspace Go (`src/`, `bin/`, `pkg/`)

### 2. Lựa Chọn Editor
- [ ] VS Code + Go extension
- [ ] GoLand (JetBrains)
- [ ] Vim/Neovim

### 3. Các Công Cụ Cơ Bản
- [ ] `go run` - Chạy file Go
- [ ] `go build` - Biên dịch chương trình
- [ ] `go mod init` - Khởi tạo module
- [ ] `go fmt` - Định dạng code
- [ ] `go vet` - Kiểm tra lỗi tiềm ẩn

---

## 📖 Phần 1: Cơ Bản

### 1.1 Cấu Trúc Chương Trình Go
- [ ] Hiểu cấu trúc file Go cơ bản
  - Package declaration
  - Imports
  - Hàm main()
  - Hàm main() và điểm bắt đầu
- [ ] Package và namespace
- [ ] Quy ước đặt tên (naming conventions)
  - Exported (chữ hoa) vs Unexported (chữ thường)

**Ví dụ:**
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

### 1.2 Biến và Kiểu Dữ Liệu
- [ ] Khai báo biến
  - `var name string`
  - `:=` short declaration
  - `const`
- [ ] Các kiểu dữ liệu cơ bản
  - **Số**: `int`, `int8`, `int32`, `int64`, `uint`, `float32`, `float64`
  - **Chuỗi**: `string`, `byte`, `rune`
  - **Boolean**: `bool`
  - **Zero values**: giá trị mặc định khi khai báo
- [ ] Type conversion
- [ ] `iota` - Sinh số tự động (cho const)

**Ví dụ:**
```go
var x int = 10
y := 20
const pi = 3.14159

// Type conversion
var a int = 100
var b float64 = float64(a)
```

### 1.3 Operators
- [ ] Arithmetic operators: `+`, `-`, `*`, `/`, `%`
- [ ] Comparison operators: `==`, `!=`, `<`, `>`, `<=`, `>=`
- [ ] Logical operators: `&&`, `||`, `!`
- [ ] Bitwise operators (nâng cao): `&`, `|`, `^`, `<<`, `>>`

### 1.4 String & Formatting
- [ ] String literals: `"..."`, `` `...` ``, `'...'`
- [ ] String manipulation
  - Concatenation (`+`)
  - Indexing và slicing
  - Raw strings
- [ ] `fmt` package
  - `Println()`, `Printf()`, `Sprintf()`
  - Format verbs: `%s`, `%d`, `%f`, `%v`, `%T`

---

## 📊 Phần 2: Cấu Trúc Dữ Liệu

### 2.1 Array
- [ ] Khai báo và khởi tạo mảng cố định
  - `var arr [5]int`
  - `arr := [3]string{"a", "b", "c"}`
- [ ] Độ dài array: `len(arr)`
- [ ] Lặp qua array: `for i, v := range arr`

**Ví dụ:**
```go
var numbers [5]int
numbers[0] = 10

for i, v := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", i, v)
}
```

### 2.2 Slice
- [ ] Hiểu slice vs array
- [ ] Khai báo slice
  - `var s []int`
  - `s := []int{1, 2, 3}`
- [ ] Slice từ array: `arr[start:end]`
- [ ] Cap và len của slice
- [ ] `append()` - Thêm phần tử
- [ ] `copy()` - Copy slice
- [ ] `make()` - Tạo slice có capacity

**Ví dụ:**
```go
s := []int{1, 2, 3}
s = append(s, 4)
newSlice := s[1:3]

s2 := make([]int, 5, 10) // len=5, cap=10
```

### 2.3 Map
- [ ] Tạo map: `m := map[string]int{"a": 1}`
- [ ] Truy cập giá trị: `m["key"]`
- [ ] Kiểm tra key tồn tại: `v, ok := m["key"]`
- [ ] Xóa phần tử: `delete(m, "key")`
- [ ] Lặp map: `for k, v := range m`
- [ ] `make()` map với capacity

**Ví dụ:**
```go
m := map[string]int{
    "apple":  5,
    "banana": 3,
}

count, exists := m["apple"]
if exists {
    fmt.Println("Có", count, "quả táo")
}
```

### 2.4 Struct
- [ ] Định nghĩa struct
- [ ] Khai báo instance
- [ ] Truy cập fields: `s.fieldName`
- [ ] Struct literals: `Person{"John", 30}`
- [ ] Embedded structs (struct trong struct)
- [ ] Struct tags

**Ví dụ:**
```go
type Person struct {
    Name string
    Age  int
    City string
}

p := Person{"John", 30, "Hanoi"}
fmt.Println(p.Name)
```

---

## ❌ Phần 3: Xử Lý Lỗi & Defer

### 3.1 Error Handling
- [ ] Interface `error`
- [ ] Return error từ hàm
- [ ] Kiểm tra error: `if err != nil`
- [ ] Tạo custom error với `errors.New()`
- [ ] Error wrapping: `fmt.Errorf()`
- [ ] `try-catch` không tồn tại trong Go

**Ví dụ:**
```go
import "errors"

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("số chia không thể bằng 0")
    }
    return a / b, nil
}

result, err := divide(10, 0)
if err != nil {
    fmt.Println("Lỗi:", err)
}
```

### 3.2 Panic & Recover
- [ ] `panic()` - Dừng chương trình
- [ ] `recover()` - Bắt panic
- [ ] Sử dụng responsibly

### 3.3 Defer
- [ ] Hiểu `defer` - Thực hiện trước khi hàm return
- [ ] LIFO (Last In First Out) order
- [ ] Cleanup resources (file, database connections)
- [ ] Defer với error handling

**Ví dụ:**
```go
func readFile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Lỗi:", err)
        return
    }
    defer file.Close() // Đóng file tự động
    
    // Đọc file...
}
```

---

## 🔗 Phần 4: Con Trỏ & Memory

### 4.1 Pointers
- [ ] Khai báo pointer: `var p *int`
- [ ] Lấy địa chỉ: `&variable`
- [ ] Dereference: `*pointer`
- [ ] Pointer để thay đổi giá trị
- [ ] Nil pointer

**Ví dụ:**
```go
x := 10
p := &x      // p là con trỏ đến x
fmt.Println(*p)  // In giá trị: 10
*p = 20      // Thay đổi giá trị thông qua con trỏ
fmt.Println(x)   // x = 20
```

### 4.2 Pointer Receivers
- [ ] Định nghĩa method với pointer receiver
- [ ] Khi nào dùng pointer vs value receiver
- [ ] Thay đổi struct thông qua pointer

**Ví dụ:**
```go
func (p *Person) setAge(age int) {
    p.Age = age
}
```

### 4.3 Memory Management
- [ ] Stack vs Heap
- [ ] Garbage Collection
- [ ] Memory leaks (ít gặp nhưng cần biết)

---

## 🔌 Phần 5: Interface & Polymorphism

### 5.1 Methods
- [ ] Định nghĩa method
- [ ] Method on struct
- [ ] Method on type bất kỳ
- [ ] Pointer vs Value receivers

**Ví dụ:**
```go
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}
```

### 5.2 Interfaces
- [ ] Định nghĩa interface
- [ ] Implicit implementation
- [ ] Interface{} (empty interface)
- [ ] Type assertion: `v, ok := x.(Type)`
- [ ] Type switch

**Ví dụ:**
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

func printArea(s Shape) {
    fmt.Println("Diện tích:", s.Area())
}
```

### 5.3 Polymorphism
- [ ] Cách Go hỗ trợ polymorphism
- [ ] Duck typing
- [ ] Interface composition

---

## ⚙️ Phần 6: Concurrency

### 6.1 Goroutines
- [ ] Khái niệm goroutine là lightweight thread
- [ ] Tạo goroutine: `go functionName()`
- [ ] Main goroutine
- [ ] Chờ goroutine hoàn thành

**Ví dụ:**
```go
go func() {
    fmt.Println("Chạy song song")
}()
```

### 6.2 Channels
- [ ] Tạo channel: `ch := make(chan int)`
- [ ] Gửi dữ liệu: `ch <- value`
- [ ] Nhận dữ liệu: `value := <-ch`
- [ ] Buffered channels
- [ ] `close()` channel
- [ ] Range trên channel

**Ví dụ:**
```go
ch := make(chan int)
go func() {
    ch <- 42
}()
value := <-ch
```

### 6.3 Synchronization
- [ ] Chủ động chờ goroutines
- [ ] `sync.WaitGroup`
- [ ] `sync.Mutex` - Lock/Unlock
- [ ] Race condition
- [ ] Deadlock

**Ví dụ:**
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // Công việc
}()
wg.Wait()
```

### 6.4 Select
- [ ] `select` statement để chờ nhiều channels
- [ ] `default` case
- [ ] Timeout patterns

---

## 📚 Phần 7: Standard Library

### 7.1 I/O & File Handling
- [ ] `os` package - File operations
  - `os.Open()`, `os.Create()`
  - `os.ReadFile()`, `os.WriteFile()`
- [ ] `io` package - Reader/Writer interfaces
- [ ] `bufio` package - Buffered I/O
- [ ] Path operations: `filepath` package

### 7.2 String & Text
- [ ] `strings` package
  - `Contains()`, `Split()`, `Join()`, `Replace()`
  - `ToUpper()`, `ToLower()`, `TrimSpace()`
- [ ] `strconv` package - Convert to/from string
  - `Atoi()`, `ParseInt()`, `Itoa()`
- [ ] Regular expressions: `regexp` package

### 7.3 Collections & Sorting
- [ ] `sort` package
  - `sort.Ints()`, `sort.Strings()`
  - `sort.Slice()` với custom comparator
- [ ] Sắp xếp struct

### 7.4 Time & Date
- [ ] `time` package
  - `time.Now()`
  - Parsing time: `time.Parse()`
  - Duration và tính toán thời gian
  - `time.Sleep()`, `time.Ticker`, `time.Timer`

**Ví dụ:**
```go
now := time.Now()
formatted := now.Format("2006-01-02 15:04:05")

duration := 5 * time.Second
time.Sleep(duration)
```

### 7.5 JSON (Rất Quan Trọng)
- [ ] Marshal: Chuyển struct sang JSON
  - `json.Marshal()`, `json.MarshalIndent()`
- [ ] Unmarshal: Chuyển JSON sang struct
  - `json.Unmarshal()`
- [ ] Struct tags: `json:"fieldName"`
- [ ] Omit empty fields: `json:"fieldName,omitempty"`

**Ví dụ:**
```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
}

data, _ := json.Marshal(user)
fmt.Println(string(data))
```

### 7.6 HTTP (Cho Web Development)
- [ ] `net/http` package
  - `http.Get()`, `http.Post()`
  - Request/Response
  - Headers
  - Status codes
- [ ] `http.Server` - Tạo server đơn giản

**Ví dụ:**
```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

http.HandleFunc("/", handler)
http.ListenAndServe(":8080", nil)
```

### 7.7 Database
- [ ] `database/sql` package
- [ ] SQL drivers (sqlite3, mysql, postgresql)
- [ ] Prepare statements
- [ ] Query và exec

---

## 🧪 Phần 8: Testing

### 8.1 Unit Testing
- [ ] Viết test file: `*_test.go`
- [ ] Test function: `func TestXxx(t *testing.T)`
- [ ] `t.Error()`, `t.Errorf()`, `t.Fatal()`
- [ ] Table-driven tests
- [ ] Assertions (sử dụng standard library)

**Ví dụ:**
```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Expected 5, got %d", result)
    }
}
```

### 8.2 Benchmarking
- [ ] `func BenchmarkXxx(b *testing.B)`
- [ ] `b.ResetTimer()`
- [ ] Chạy benchmark: `go test -bench=.`

### 8.3 Test Coverage
- [ ] `go test -cover`
- [ ] HTML coverage report

---

## 💡 Phần 9: Thực Hành Dự Án

### Dự Án Mức 1: Beginner
1. **Command-line Calculator**
   - Đọc input từ user
   - Thực hiện các phép toán cơ bản
   - Xử lý lỗi

2. **TODO List Application**
   - CRUD operations
   - Lưu vào file JSON
   - Command-line interface

3. **File Searcher**
   - Đọc từ directory
   - Tìm kiếm file theo tên
   - In kết quả

### Dự Án Mức 2: Intermediate
1. **Web Server đơn giản**
   - Xử lý HTTP requests
   - Serve static files
   - Routing cơ bản

2. **REST API**
   - CRUD endpoints
   - JSON request/response
   - Error handling

3. **Web Scraper**
   - Fetch HTML
   - Parse với regex hoặc HTML parser
   - Process data

### Dự Án Mức 3: Advanced
1. **Chat Application**
   - Goroutines cho connections
   - Channels để gửi messages
   - Broadcasting

2. **Database Application**
   - SQL operations
   - ORM (gorm)
   - Connection pooling

3. **Microservice**
   - Multiple services
   - Communication (HTTP/gRPC)
   - Configuration management

---

## 📌 Chú Ý Quan Trọng

- ✅ Go là statically typed - lỗi sẽ bị phát hiện tại compile time
- ✅ Go force formatting - sử dụng `go fmt`
- ✅ Go không có try-catch - sử dụng explicit error checking
- ✅ Concurrency là điểm mạnh của Go - bắt đầu sử dụng goroutines sớm
- ✅ Interface rất flexible - tận dụng duck typing
- ✅ Standard library rất mạnh mẽ - ưu tiên sử dụng trước third-party packages
- ✅ Dependency management với `go.mod` - hiểu cách quản lý version

---

## 🔗 Tài Nguyên Hữu Ích

- [Go Official Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go) - Best practices
- [Go by Example](https://gobyexample.com/)
- [Learn Go with Tests](https://github.com/quii/learn-go-with-tests)
- [Go Language Specification](https://golang.org/ref/spec)

---

## 📝 Hướng Dẫn Sử Dụng Roadmap Này

1. Đọc từ trên xuống dưới theo thứ tự
2. Với mỗi chủ đề, hãy:
   - [ ] Đọc khái niệm
   - [ ] Viết code ví dụ
   - [ ] Thử modification nhỏ
   - [ ] Hoàn thành các bài tập
3. Khi xong một phần, hãy dùng những kiến thức đó trong dự án nhỏ
4. Đừng đi tiếp trước khi nắm vững phần hiện tại
5. Revisit các phần khó nhất 2-3 lần

---

**Chúc bạn học tập thành công! 🎯**
