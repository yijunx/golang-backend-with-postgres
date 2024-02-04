build:
	@go build -o bin/gobackend

up: build
	@./bin/gobackend

test:
	@go test -v ./...