#!/usr/bin/env bash

docker run -it -v $(pwd):/go/src/HFAT golang:1.7.3 /bin/bash /go/src/HFAT/ci/test.sh