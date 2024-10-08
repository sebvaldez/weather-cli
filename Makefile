# Define variables
BINARY_NAME=weather
BUILD_DIR=bin

# Default target
all: build

# Build the binary for the current OS
build:
	@echo "Building the binary for the current OS..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

# Build binaries for multiple platforms
build-all:
	@echo "Building binaries for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 .
	# GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .
	# GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .
	# GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 .
	# GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe .

# Run tests with colorized output using gotestsum
test-internals:
	@echo "Running tests..."
	go run gotest.tools/gotestsum --format testname -- ./internal/...

# Clean up generated files
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

# Run the application with arguments
run:
	go run main.go $(filter-out $@,$(MAKECMDGOALS))

# Prevent make from interpreting the arguments as Makefile targets
%:
	@:

# Print help
help:
	@echo "Makefile commands:"
	@echo "  build      - Build the project for the current platform"
	@echo "  build-all  - Build the project for multiple platforms"
	@echo "  test       - Run tests"
	@echo "  clean      - Clean up generated files"
	@echo "  run        - Run the application"
	@echo "  help       - Show this help message"

.PHONY: all build build-all test clean run help
