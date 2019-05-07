#!/usr/bin/env bash

# -tags netgo apline构建golang编译问题
GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -tags netgo -o main

# docker build
docker build -f ./Dockerfile -t registry.cn-hangzhou.aliyuncs.com/dreamlu/common:deercoder-chat-api-gateway .

# remove build
rm -rf main