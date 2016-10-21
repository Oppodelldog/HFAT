#!/usr/bin/env bash

echo "lint & vet for compiler warnings"
go get -u github.com/golang/lint/golint

golint HFAT/... > /go/src/HFAT/lint.txt
go tool vet src/HFAT > /go/src/HFAT/vet.txt