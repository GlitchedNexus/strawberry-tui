.PHONY: build test run fmt lint

build:
	go build ./...

test:
	go test ./...

run:
	go run ./cmd/tui-playground

fmt:
	gofmt -w .

