run: build
	@./bin/msgq

build:
	@go build -o bin/msgq

test:
	@go test -v ./...
