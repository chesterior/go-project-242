BINARY_NAME=hexlet-path-size

setup:
	go mod tidy

build:
	go build -o bin/$(BINARY_NAME) ./cmd/$(BINARY_NAME)

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

test:
	go test -v ./...