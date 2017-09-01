#!/bin/bash
GOOS=linux go build -o build/server/user_server_linux_amd64 server/user/server.go && GOOS=linux go build -o build/server/time_server_linux_amd64 server/time/server.go && GOOS=linux go build -o build/server/log_server_linux_amd64 server/log/server.go server/log/writer.go
