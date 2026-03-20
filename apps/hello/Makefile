## Build the application
## Usage: make build
.PHONY: build
build:
	go build -o bin/hello .

## Run the application
## Usage: make run
.PHONY: run
run:
	./bin/hello

## Install the application
## Usage: make install
.PHONY: install
install:
	go install .

## Uninstall the application
## Usage: make uninstall
## Note: $$ trong Makefile → $ trong shell (để dùng biến môi trường GOBIN).
.PHONY: uninstall
uninstall:
	@if [ -n "$$GOBIN" ]; then rm -f "$$GOBIN/hello"; fi
	rm -f "$$(go env GOPATH)/bin/hello"

## Run the tests
## Usage: make test
.PHONY: test
test:
	go test -v .

## Clean the build artifacts
## Usage: make clean
.PHONY: clean
clean:
	rm -rf bin/**