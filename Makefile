#!/bin/bash

test:
	go test -v ./...

build:
	go build -o bin/ ./...