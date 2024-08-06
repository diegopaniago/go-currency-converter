#!/bin/bash

test:
	go test -v ./...

buid:
	go build -o bin/ ./...