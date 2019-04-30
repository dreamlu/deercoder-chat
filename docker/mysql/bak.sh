#!/bin/bash
# crontab :1 2 * * * xxx/bak.sh
# docker备份目录/不同于脚本执行目录/xxx/bak.sh
# 覆盖问题
echo "进入xxx目录进行备份"
cd /bak/
DATE=`date +%Y%m%d`
OLD=`date -d"1 day ago" +"%Y%m%d"`
OLD3=`date -d"3 day ago" +"%Y%m%d"`
USER="${2}"
PASSWORD="${3}"
DATABASE="${1}"

mysqldump -R ${DATABASE} -u${USER} -p${PASSWORD} | gzip > ${1}_${DATE}.sql.gz
if tar zcf ${1}_${DATE}.tar.gz ${1}_${DATE}.sql.gz; then
	rm ${1}_${OLD3}.tar.gz
	rm ${1}_${OLD}.sql.gz
fi
