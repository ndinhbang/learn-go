module github.com/ndinhbang/learn-go/apps/hello

go 1.26.1

// go mod edit -replace github.com/ndinhbang/learn-go/pkg/greetings=../../pkg/greetings
replace github.com/ndinhbang/learn-go/pkg/greetings => ../../pkg/greetings

require github.com/ndinhbang/learn-go/pkg/greetings v0.0.0-00010101000000-000000000000
