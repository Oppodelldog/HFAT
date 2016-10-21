#!/usr/bin/env bash

go get -u github.com/golang/lint/golint

# Run lint tools (Compiler warning plugin)
golint HFAT/... > lint.txt

# Run vet tools (Compiler warning plugin)
go vet HFAT/... > vet.txt