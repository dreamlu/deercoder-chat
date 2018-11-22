#!/bin/bash
#ssh no secret run cmd
#up.sh made by lucheng
#version 1.1
#server ip
ip=47.94.8.188
uploadDir=root
#directory
curDir=$(basename `pwd`)
#get the newest go build processname
GOOS=linux GOARCH=amd64 go build -o $1
echo "compress tar.gz is running..."
tar -zcf $1.tar.gz $1
echo "$1 compile finished"
echo "now upload the file"
scp $1.tar.gz root@$ip:/$uploadDir/$1
echo "$1 upload finshed"
echo "and now restart the programe"
ssh root@$ip

echo "upload and run over,remove $1.tar.gz"
rm -rf $1.tar.gz
#echo "2s later exit"
#sleep 2
#exit
