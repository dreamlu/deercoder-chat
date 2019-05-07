#!/usr/bin/env bash
GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -tags netgo -o main

# docker build
docker build -f ./Dockerfile -t registry.cn-hangzhou.aliyuncs.com/dreamlu/common:deercoder-chat-user-srv .

# remove build
rm -rf main