#!/usr/bin/env bash

# -tags netgo apline构建golang编译问题

# go mod 中的静态资源引入问题
#GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main
GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -tags netgo -o main

# docker build
docker build -f ./Dockerfile -t registry.cn-hangzhou.aliyuncs.com/dreamlu/common:deercoder-chat-api .

# remove build
rm -rf main