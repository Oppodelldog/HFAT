#!/usr/bin/env bash

echo "lint & vet for compiler warnings"
go get -u github.com/golang/lint/golint

golint HFAT/... | sed s/.*HFAT/\./g > /go/src/HFAT/lint.txt
