APP_NAME := go-be
BIN := bin/$(APP_NAME)

.PHONY: run build clean test fmt tidy

run:
	go run ./cmd/api

build:
	mkdir -p bin
	go build -o $(BIN) ./cmd/api

clean:
	rm -rf bin

test:
	go test ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy


