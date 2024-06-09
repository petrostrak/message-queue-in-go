run: build
	@./bin/msgq

build:
	@go build -o bin/msgq
