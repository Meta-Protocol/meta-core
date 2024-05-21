#!/usr/bin/env bash

# Install golines
go install github.com/segmentio/golines@v0.12.2

# Run golines in Cosmos modules and ZetaClient codebase
find . -type f -name '*.go' -exec golines -w --max-len=120 {} + > /dev/null 2>&1

# Print a message to indicate completion
echo "Go source code lines formatted."