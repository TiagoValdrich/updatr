
install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.7 &&\
	go mod tidy && \
	go mod vendor

run:
	go run main.go

build:
	go build -o updatr main.go

lint:
	golangci-lint run ./...
