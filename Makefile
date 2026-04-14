setup:
	cd code && go mod tidy

build:
	cd code && go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

lint:
	cd code && golangci-lint run

lint-fix:
	cd code && golangci-lint run --fix

test:
	cd code && go test -v ./...