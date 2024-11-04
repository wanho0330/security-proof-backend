#!/bin/bash

# Run golangci-lint with specified configuration
golangci-lint run ../... -c /Users/wanho/GolandProjects/security-proof/.golangci.yml

echo "GolangCI-Lint completed with configuration from .golangci.yml."
