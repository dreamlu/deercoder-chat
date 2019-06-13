#!/bin/bash
# crontab :1 2 * * * xxx/bak.sh
# mysql 备份
# docker/本机备份两种方式
# 详情./bak.sh -h查看
DATE=`date +%Y%m%d`
OLD=`date -d"1 day ago" +"%Y%m%d"`
OLD3=`date -d"3 day ago" +"%Y%m%d"`
# 备份文件目录
bakDir=
# 数据信息
USER=
PASSWORD=
DATABASE=
# 容器名或id
CONTAINER=

# 参数处理
# :需要参数
while getopts ":c:u:p:d:b:h" opt
do
    case ${opt} in
        c)
        CONTAINER=$OPTARG
        echo "参数CONTAINER的值${CONTAINER}"
        ;;
        u)
        USER=$OPTARG
        echo "参数USER的值${USER}"
        ;;
        p)
        PASSWORD=$OPTARG
        echo "参数PASSWORD的值${PASSWORD}"
        ;;
        d)
        DATABASE=$OPTARG
        echo "参数DATABASE的值${DATABASE}"
        ;;
        b)
        bakDir=$OPTARG
        echo "参数bakDir的值${bakDir}"
        ;;
        h)
        echo -e "-c 容器名或id\n-u 数据库账号名\n-p 数据库密码\n-b 备份目录\n-d 备份的数据库名"
        exit 1
        ;;
        ?)
        echo "未知参数"
        exit 1;;
    esac
done

# 参数验证
echo -e "数据库: \n用户名${USER}\n密码${PASSWORD}\n数据库名${DATABASE}"
echo -e "容器名: ${CONTAINER}\n"
echo -e "备份文件目录: ${bakDir}\n"

# 开始备份
echo "备份类型："
if [[ -n ${CONTAINER} ]]; then
    echo "容器备份"
    docker exec ${CONTAINER} /usr/bin/mysqldump -u ${USER} --password=${PASSWORD} ${DATABASE} | gzip > ${bakDir}/${DATABASE}_${DATE}.sql.gz
else
    echo "本机备份"
    mysqldump -u ${USER} -p ${PASSWORD} -R ${DATABASE} | gzip > ${bakDir}/${DATABASE}_${DATE}.sql.gz
fi
#docker exec ${CONTAINER} /usr/bin/mysqldump -u ${USER} --password=${PASSWORD} ${DATABASE} | gzip > ${bakDir}/${DATABASE}_${DATE}.sql.gz
if tar zcf ${bakDir}/${DATABASE}_${DATE}.tar.gz ${bakDir}/${DATABASE}_${DATE}.sql.gz; then
	rm ${bakDir}/${DATABASE}_${OLD3}.tar.gz
	rm ${bakDir}/${DATABASE}_${OLD}.sql.gz
fi
