#!/bin/bash
env GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC=musl-gcc go build -o tld-go-linux-musl-amd64 main.go