#!/bin/bash

# Build for Linux arm64
CC="aarch64-unknown-linux-gnu-gcc" CXX="aarch64-unknown-linux-gnu-g++" CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o bin/snowguard-linux-arm64 main.go

# Build for OSX
GOOS=darwin GOARCH=arm64 go build -o bin/snowguard-darwin-arm64 main.go
