#!/bin/bash

VERSION="1.0"
COMMIT="abcdefghijklmnopg"
TARGET="main/main.go"

go build  -gcflags "-N -l" -ldflags "-X other.VERSION=1.0.0 -X gpxj/other.VERSION=1.0.3 -X main.VERSION=1.0.0 -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'"  -v -o "server-${VERSION}-${COMMIT:0:7}" $TARGET 

