websocket IM  

微服务-即时聊天/设计  

- 特点  
1.内部rpc  
2.对外json  
3.工具类：[deercoder-gin](https://github.com/dreamlu/deercoder-gin)  
4.websocket即时通讯  
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

- 好友列表(will like this)
![好友列表](common-srv/static/img/chat.png)  

- 添加好友  
![添加好友](common-srv/static/img/addFriend.png)  

- 开发/部署(待优化)  
1.开发模式修改执行[devMode.sh](devMode.sh)文件    
2.本地直接进行开发即可  
3.docker部署先执行每个模块下docker.sh命令  
4.docker目录下pushAll.sh脚本推向私有或共有仓库  
5.服务器上执行docker目录下pullAll.sh脚本, 拉取镜像  
6.docker-compose/k8s部署参考docker目录  
