websocket IM  

微服务-即时聊天/设计[RPC]  

[rpc版本, 理应一个服务提供rpc和http两种接口, 这里http接口统一在`api`模块中, `api-gateway`不是很有必要]

- 特点  
1.内部rpc  
2.对外json  
3.工具类：[go-tool](https://github.com/dreamlu/go-tool)  
4.websocket即时通讯[[核心代码](./api/controllers/chat/chatWs.go)]  
5.[在线接口查看](https://www.eolinker.com/#/share/index?shareCode=jH4ISr)  
6.[go micro](https://github.com/micro/go-micro) docker化, 具体见Dockerfile  
7.[docker-compose支持](docker/docker-compose.yaml)   
8.[k8s支持](docker/k8s/README.md)  
n....  

> 状态-ing  

- 扩展  
1.mysql自动备份,[备份规则](docker/mysql/bak.sh)  
2....  

- 微服务各个模块运行状态(这里consul进行服务发现)  
[http://localhost:8500](http://localhost:8500)

- 登录  
![登录页面](common-srv/static/img/login.png)  

- 好友列表以及聊天(图片上传)  
![好友列表](common-srv/static/img/fileUpLoad.png)  

- 添加好友  
![添加好友](common-srv/static/img/addFriend.png)  

- 开发/部署  
执行[devMode.sh](devMode.sh)脚本,修改开发模式  

- 本地开发  
1.运行consul服务发现    
2.各个模块下`./run.sh `    

- docker部署  
1.各个模块下`./docker.sh`打包镜像  
2.docker目录下pushAll.sh脚本推向私有或共有仓库  
3.服务器上执行docker目录下pullAll.sh脚本, 拉取镜像  
4.docker-compose/k8s部署参考docker目录  
