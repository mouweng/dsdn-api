#!/bin/sh
# export GOROOT=/usr/local/go
printf "\033[1;33m"
echo "Begin to build......"
go build -o main main.go
echo "build finished"
printf "\033[m"
