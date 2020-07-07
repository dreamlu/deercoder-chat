#!/usr/bin/env bash
# 运行docker-compose, 启动程序
docker-compose up -d
# 安装备份脚本
#./mysql/setup.sh -c 容器名或id -u 数据库账号 -p 数据库密码 -d 备份的数据库名 -b 备份目录(/root/bak)
./mysql/setup.sh -c w-mysql -u dreamlu -p xxxx -d xxx -b /root/bak
