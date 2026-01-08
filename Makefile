.PHONY: build install clean test build-all

# Variables
BINARY_NAME=gwt
VERSION=$(shell git describe --tags --always 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"
GO=go
GOFLAGS=-v

# Build directory
BUILD_DIR=dist

# Default target
.DEFAULT_GOAL := help

help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'

build: ## Build the binary for current platform
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@$(GO) build $(GOFLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) main.go
	@echo "Built: $(BUILD_DIR)/$(BINARY_NAME)"

install: build ## Install the binary to $$GOPATH/bin
	@echo "Installing $(BINARY_NAME)..."
	@$(GO) install $(LDFLAGS) .
	@echo "Installed to: $$(go env GOBIN)/$(BINARY_NAME)"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@$(GO) clean
	@echo "Clean complete"

test: ## Run tests
	@echo "Running tests..."
	@$(GO) test -v ./...

build-all: ## Build for all platforms
	@echo "Building for all platforms..."
	@mkdir -p $(BUILD_DIR)
	@echo "Building for Linux (amd64)..."
	@GOOS=linux GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 main.go
	@echo "Building for Linux (arm64)..."
	@GOOS=linux GOARCH=arm64 $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 main.go
	@echo "Building for macOS (amd64)..."
	@GOOS=darwin GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 main.go
	@echo "Building for macOS (arm64)..."
	@GOOS=darwin GOARCH=arm64 $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 main.go
	@echo "Building for Windows (amd64)..."
	@GOOS=windows GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe main.go
	@echo "Build complete. Binaries in: $(BUILD_DIR)/"
	@ls -lh $(BUILD_DIR)/

fmt: ## Format code
	@echo "Formatting code..."
	@$(GO) fmt ./...
	@echo "Format complete"

lint: ## Run linter
	@echo "Running linter..."
	@$(GO) vet ./...
	@echo "Lint complete"

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@$(GO) mod download
	@$(GO) mod tidy
	@echo "Dependencies downloaded"

version: ## Show version
	@echo "$(BINARY_NAME) version: $(VERSION)"
