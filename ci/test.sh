#!/usr/bin/env bash

echo "STARTING GO TESTS"
echo "get go-junit-report"
go get -u github.com/jstemmer/go-junit-report

echo "run tests"
go test -v ./... | go-junit-report > /go/src/HFAT/report.xml