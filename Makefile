# Makefile for Monte Carlo Pi Simulation

# Variables
BINARY_NAME=pi
GO=go

# Default target: show help
.PHONY: help
help:
	@echo "Monte Carlo Pi Simulation Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make build    - Build the application"
	@echo "  make run      - Run the application"
	@echo "  make clean    - Remove build artifacts"
	@echo "  make help     - Show this help message"

# Build the application
.PHONY: build
build:
	$(GO) build -o $(BINARY_NAME) main.go

# Run the application
.PHONY: run
run: build
	./$(BINARY_NAME)

# Clean build artifacts
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)