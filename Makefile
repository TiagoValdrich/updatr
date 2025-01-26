
install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2 \
	go mod tidy && \
	go mod vendor

run:
	go run cmd/main.go

build:
	go build cmd/main.go

lint:
	golangci-lint run ./...
