# Makefile for building, installing, and cleaning spiderNetscan

# Define binary name and the main Go file
BINARY_NAME=spiderNetscan
MAIN_FILE=cmd/spiderNetscan.go

# Build the binary
build:
	@echo "Building the binary..."
	go build -o $(BINARY_NAME) $(MAIN_FILE)

# Install the binary to /usr/local/bin
install: build
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	sudo mv $(BINARY_NAME) /usr/local/bin/

# Clean up the built binary
clean:
	@echo "Cleaning up the binary..."
	rm -f $(BINARY_NAME)

# Update the tool from GitHub 
update:
	@echo "Updating spiderNetscan from GitHub..."
	git pull origin main

# Print help text for make commands
help:
	@echo "Makefile Commands:"
	@echo "  build    - Build the spiderNetscan tool"
	@echo "  install  - Install the tool to /usr/local/bin"
	@echo "  clean    - Remove the built binary"
	@echo "  update   - Update the tool from GitHub"
	@echo "  help     - Display this help text"
