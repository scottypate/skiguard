#!/bin/bash

# Build for Linux arm64
CC="aarch64-unknown-linux-gnu-gcc" CXX="aarch64-unknown-linux-gnu-g++" CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o bin/skiguard-linux-arm64 main.go

# Build for Linux amd64
CC="x86_64-linux-gnu-gcc" CXX="x86_64-linux-gnu-g++" CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o bin/skiguard-linux-amd64 main.go

# Build for OSX
GOOS=darwin GOARCH=arm64 go build -o bin/skiguard-darwin-arm64 main.go
