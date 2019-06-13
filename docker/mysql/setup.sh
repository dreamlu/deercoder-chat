#!/bin/bash
if [[ $# -lt 1 ]]; then
    	echo "没有命令行参数,请重新执行"
	exit
fi
# 定时备份文件
cronFile="/etc/cron.d/crontab"

if [[ ! -f "$cronFile" ]]; then
  	touch ${cronFile}
else
  	echo "定时任务文件${cronFile}已存在"
fi
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

if grep "bak.sh" ${cronFile} ; then
	echo "自动备份脚本已存在!"
else
   	# 将命令写入定时备份文件
   	# 每天2:01开始备份数据
	echo -e "#定时备份数据至${bakDir}目录\n1 2 * * * root ${bakDir}/bak.sh -c ${CONTAINER} -u ${USER} -p ${PASSWORD} -d ${DATABASE} -b ${bakDir}" >> ${cronFile}
	echo "定时备份数==> ${bakDir}中"
fi
