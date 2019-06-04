ARG MYSQL_VERSION

FROM mysql:$MYSQL_VERSION

# 参数问题:https://stackoverflow.com/questions/48831447/docker-compose-build-args-not-passing-to-dockerfile
ARG MYSQL_DATABASE
ARG MYSQL_USER
ARG MYSQL_PASSWORD

# 工作目录不可与挂载目录相同, 否则复制进取的文件, 会被宿主机挂载目录覆盖
WORKDIR /$MYSQL_DATABASE/bak/
COPY ["setup.sh","bak.sh","deercoder-chat.sql","./"]

# 导入sql文件
# 待完善
#RUN mysql -u$MYSQL_USER -p$MYSQL_PASSWORD deercoder-chat < deercoder-chat.sql

# 更换源,apt update 卡住!!?
#RUN sed -i "s/http:\/\/repo.mysql.com\/apt\/debian/https:\/\/mirrors.aliyun.com\/apt\/debian\//g" /etc/apt/sources.list.d/mysql.list
#RUN sed -i "s/http:\/\/deb.debian.org\/debian/https:\/\/mirrors.aliyun.com\/debian\//g" /etc/apt/sources.list
#RUN sed -i "s/http:\/\/security.debian.org\/debian-security/https:\/\/mirrors.aliyun.com\/debian-security\//g" /etc/apt/sources.list

# 定时任务软件包
# nohup 可后台执行
RUN nohup apt update && apt install cron -y --no-install-recommends &
# 定时任务设置
RUN chmod 755 setup.sh bak.sh
RUN ./setup.sh "$MYSQL_DATABASE" "$MYSQL_USER" "$MYSQL_PASSWORD"
EXPOSE 3306