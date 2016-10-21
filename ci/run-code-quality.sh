#!/usr/bin/env bash

docker run -v $(pwd):/go/src/HFAT golang:1.7.3 /bin/bash /go/src/HFAT/ci/docker-container/test-junit.sh