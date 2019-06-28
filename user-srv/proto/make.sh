#!/usr/bin/env bash
protoc --proto_path=${GOPATH}/gowork/work:. --micro_out=. --go_out=. *.proto