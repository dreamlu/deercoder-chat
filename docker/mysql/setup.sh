#!/bin/bash
if [[ $# -lt 1 ]]; then
    	echo "没有命令行参数,请重新执行"
	exit
fi
bakdir="/etc/cron.d/crontab"
if [[ ! -f "$bakdir" ]]; then
  	touch ${bakdir}
else
  	echo "定时任务文件${bakdir}已存在"
fi

if grep "bak.sh" $bakdir ; then
	echo "自动备份脚本已存在!"
else
    # 每天2:01开始备份数据
	echo -e "#定时备份${1}数据\n1 2 * * * root /$1/bak/bak.sh $1 $2 $3" >> $bakdir
	echo "定时备份==> ${1} 中的数据"
fi
