
build:
	@echo "Building the project..."
	@go build -o bin/ ./cmd/faker/

run:
	@echo "Running the project..."
	@go run ./cmd/faker/

test:
	@echo "Running tests..."
	@go test -v ./...

clean:
	@echo "Cleaning up..."
	@rm -rf bin/

.PHONY: build run test clean


