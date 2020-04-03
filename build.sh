#!/bin/env bash

mkdir -p ./bin

# build static binaries for each architecture with debugging info stripped
CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/krucible-linux-amd64  main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/krucible.exe          main.go
CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/krucible-darwin-amd64 main.go
