BINARY_NAME=spiderNetscan
MAIN_FILE=cmd/spiderNetscan.go
VERSION=v1.0.1

build:
	@echo "Building the binary (Version: $(VERSION))..."
	go build -ldflags "-X 'main.version=$(VERSION)'" -o $(BINARY_NAME) $(MAIN_FILE)

install: build
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	sudo mv $(BINARY_NAME) /usr/local/bin/

clean:
	@echo "Cleaning up the binary..."
	rm -f $(BINARY_NAME)

update:
	@echo "Updating spiderNetscan from GitHub..."
	git pull origin main

help:
	@echo "Makefile Commands:"
	@echo "  build    - Build the spiderNetscan tool with version info"
	@echo "  install  - Install the tool to /usr/local/bin"
	@echo "  clean    - Remove the built binary"
	@echo "  update   - Update the tool from GitHub"
	@echo "  help     - Display this help text"
