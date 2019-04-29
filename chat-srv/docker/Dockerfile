FROM alpine
MAINTAINER dreamlu <862362681@qq.com>

WORKDIR /app
#COPY conf \
#     static \
#     yyd \
#     /root/yyd/
COPY ../conf /app/conf
COPY ../main /app/

CMD ["./main"]

#更新Alpine的软件源为国内（清华大学）站点
RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.4/main/" > /etc/apk/repositories

RUN apk update \
        && apk upgrade \
        && apk add --no-cache bash \
        bash-doc \
        bash-completion \
        && rm -rf /var/cache/apk/* \
        && /bin/bash