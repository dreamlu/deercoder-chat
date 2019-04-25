#!/usr/bin/env bash
protoc --proto_path=$GOPATH/gowork/src:. --micro_out=. --go_out=. *.proto