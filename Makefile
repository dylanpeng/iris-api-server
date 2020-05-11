#!/bin/bash
export LANG=zh_CN.UTF-8

ENVARG=GOPATH=$(CURDIR) GO111MODULE=on
LINUXARG=CGO_ENABLED=1 GOOS=linux GOARCH=amd64
BUILDARG=-ldflags " -s -X main.buildTime=`date '+%Y-%m-%dT%H:%M:%S'` -X main.gitHash=(`git symbolic-ref --short -q HEAD`)`git rev-parse HEAD`"

export GOPATH

dep:
	cd src; ${ENVARG} go get ./...; cd -

gateway:
	cd src/gateway; ${ENVARG} go build ${BUILDARG} -o ../../bin/gateway main.go;