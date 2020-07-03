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

p:
	mkdir -p src/lib/proto
	rm -rf src/lib/proto/*

	cd src; protoc -I ../protocol/juggernaut --gofast_out=plugins=grpc:. common/base.proto; cd -
	cd src; protoc -I ../protocol/juggernaut --gofast_out=plugins=grpc:. common/message.proto; cd -
	cd src; protoc -I ../protocol/juggernaut --gofast_out=plugins=grpc:. service/gateway.proto; cd -

	mv src/juggernaut/lib/api-proto src/lib/proto/juggernaut

	ls src/lib/proto/*/*/*/*.pb.go | xargs sed -i -e "s/,omitempty//"
	ls src/lib/proto/juggernaut/*/*/*.pb.go | xargs sed -i -e "s@api-proto/@proto/juggernaut/@"
	find src/lib/proto -name "*.pb.go-e"  | xargs rm -f

	rm -rf src/juggernaut

clean:
	rm -fr bin/*
	rm -fr lbin/*
	chmod -R 766 pkg/*
	\rm -r pkg/*
	rm src/go.sum