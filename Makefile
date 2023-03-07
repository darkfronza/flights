.DEFAULT_GOAL := build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o app

test:
	go test -v ./...

.PHONY: fmt vet build test