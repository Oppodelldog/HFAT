#!/usr/bin/env bash
echo "getting tools for code cover report conversion (for cobertura plugin)"
go get github.com/axw/gocov/...
go get github.com/AlekSi/gocov-xml
gocov test ./... | gocov-xml > /go/src/HFAT/coverage.xml
