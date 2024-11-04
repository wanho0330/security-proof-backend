#!/bin/bash

# Set GOPRIVATE environment variable
export GOPRIVATE=buf.build/gen/go

# Run go get commands to fetch the packages
go get buf.build/gen/go/wanho/security-proof-api/connectrpc/go@latest
go get buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go@latest

echo "Dependencies have been successfully fetched."
