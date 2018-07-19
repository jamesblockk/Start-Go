# Go - gRPC

## Install

$ go get -u google.golang.org/grpc
$ brew install protobuf
$ go get github.com/golang/protobuf/protoc-gen-go
$ export PATH=$PATH:$GOPATH/bin

## Setup

in pb
$ protoc --go_out=plugins=grpc:. *.proto

### reference

https://yami.io/grpc/