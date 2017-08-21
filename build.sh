#!/bin/bash
GOOS=darwin go build -o build/client_darwin_amd64 client.go
GOOS=darwin go build -o build/server_darwin_amd64 server.go
GOOS=linux go build -o build/client_linux_amd64 client.go
GOOS=linux go build -o build/server_linux_amd64 server.go
