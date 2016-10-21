#!/usr/bin/env bash

echo "getting tool for test report conversion (for junit plugin)"
go get -u github.com/jstemmer/go-junit-report

echo "run tests"
go test -v ./... | go-junit-report > /go/src/HFAT/report.xml