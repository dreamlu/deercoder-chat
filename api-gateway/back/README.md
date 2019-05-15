后台配置  

- 界面  

![后台界面](../../common-srv/static/img/bank.png)  

- 地址  
http://域名/admin/login(暂不支持https)  

来源:https://github.com/chenhg5/go-admin  

- 配置  
默认后台关闭了,em...因有不少问题  
开启方式, 全局搜索 `//back.SetBack(router)`, 取消注释即可  
另外需要导入后台配置数据库,当前目录下的admin.sql  