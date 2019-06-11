#!/usr/bin/env bash

# 本地免consul安装
# docker安装方式
# 这里consul 53端口占用
docker run -p 8300:8300 -p 8301:8301 -p 8301:8301/udp -p 8302:8302/udp -p 8302:8302 -p 8400:8400 -p 8500:8500 -p 54:53/udp --name=d-consul consul

# remove
docker rm d-consul