#!/bin/bash

# Ensure that the script is run from the project root directory
if [ ! -f go.mod ]; then
  echo "This script must be run from the project root directory!"
  exit 1
fi

# Install Go dependencies
echo "Installing Go dependencies..."
go mod tidy

# Build the tool
echo "Building the tool..."
make build

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "Build successful!"
else
  echo "Build failed. Please check the error messages above."
  exit 1
fi

# Move the binary to /usr/local/bin for global access
echo "Installing spiderNetscan..."
sudo mv spiderNetscan /usr/local/bin/

# Confirm installation
if command -v spiderNetscan &>/dev/null; then
  echo "Installation successful! You can now run 'spiderNetscan' from anywhere."
else
  echo "Installation failed. Please try again."
  exit 1
fi

# Additional instructions
echo "To update spiderNetscan in the future, you can run:"
echo "'spiderNetscan --update'"

echo "Installation complete!"
