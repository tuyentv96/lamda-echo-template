.DEFAULT_GOAL := all

.PHONY: default test

build:
    env GOOS=linux go build -ldflags="-s -w" -o bin/echo-lamda cmd/main.go