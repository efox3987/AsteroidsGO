default:
	build
	libs
	run

build:
	@echo "Building..."
	@go build -o bin/ ./cmd/...

libs:
	@echo "Getting dependencies..."
	@go mod download
	@go mod tidy

run:
	@echo "Running..."
	@./bin/$(shell basename $(CURDIR))
