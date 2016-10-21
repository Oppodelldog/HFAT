#!/usr/bin/env bash

docker run -v "$(pwd)":/go/src/HFAT golang:1.7.3 /bin/bash -c "src/HFAT/ci/docker-container/test-junit.sh && src/HFAT/ci/docker-container/test-cobertura.sh  && src/HFAT/ci/docker-container/compiler-warnings.sh"