#### 镜像拉取  

- 镜像拉取  
1.执行`./pullAll.sh`拉取共有或私有镜像  
n...  

#### docker/docker-compose支持  

- docker测试  
1.修改conf下app.conf中devMode = dev 为prod  
2.运行上级目录下docker.sh  
3.docker-compose 运行(问题参考下面)   

- 问题  
1.启动顺序待完善  
2.docker-compose up -d mysql,后台运行  
3.docker-compose up demo, 运行demo, 查看请求结果  
4.other待完善  