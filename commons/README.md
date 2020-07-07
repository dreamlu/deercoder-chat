# common 工具类

- 常用命令  
1. protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. ./*.proto  
2.docker run -p 8300:8300 -p 8301:8301 -p 8301:8301/udp -p 8302:8302/udp -p 8302:8302 -p 8400:8400 -p 8500:8500 -p 54:53/udp consul