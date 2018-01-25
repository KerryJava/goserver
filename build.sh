#!/bin/bash

VERSION="1.0"
TARGET="main/main.go"
PREFIX="github.com/KerryJava/goserver"
echo "GOPATH is $GOPATH"
echo "git 状态"
git log --stat -n 1

COMMIT=$(git log | head -n 1 | cut -d ' ' -f 2)
echo $COMMIT

FLAGS=" -X '${PREFIX}/other.VERSION=${VERSION}' -X '${PREFIX}/other.BUILD_TIME=`date`' -X '${PREFIX}/other.GO_VERSION=`go version`' -X '${PREFIX}/other.COMMIT=${COMMIT}' \
  	-X main.VERSION=${VERSION} -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'"
#FLAGS="-X main.VERSION=1.0.0"

go build -gcflags "-N -l" \
       	-ldflags "${FLAGS}" \
	-v \
	-o "server-${VERSION}-${COMMIT:0:7}" \
	$TARGET 

