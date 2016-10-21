#!/usr/bin/env bash

echo "STARTING GO TESTS"

echo "getting tool for test report conversion (for junit plugin)"
go get -u github.com/jstemmer/go-junit-report

echo "getting tools for code cover report conversion (for cobertura plugin)"
go get github.com/axw/gocov/...
go get github.com/AlekSi/gocov-xml
gocov test ./... | gocov-xml > /go/src/HFAT/coverage.xml

echo "run tests"
go test -v ./... | go-junit-report > /go/src/HFAT/report.xml