build:
	@go build -o bin/gobackend

run: build
	@./bin/gobackend

test:
	@go test -v ./...