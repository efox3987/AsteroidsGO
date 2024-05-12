default: libs build run

build:
	@echo "Building..."
	@go build -o bin/ ./...

libs:
	@echo "Getting dependencies..."
	@go mod download
	@go mod tidy

run:
	@echo "Running..."
	@./bin/$(shell basename $(CURDIR))

clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@go clean

all: clean default

