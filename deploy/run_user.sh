#!/bin/bash
# please run with root or sudo
# create user for run 

# 用户名密码
worker="buaashow"
password=$(perl -e 'print crypt($ARGV[0], "password")' "rott@123")
# 运行脚本
runsh=$PWD/run.sh

# 判断用户是否存在
user_exists() {
    local ret='0'
    id -u $1 >/dev/null 2>&1 || { local ret='1'; }

    if [ "$ret" -ne 0 ]; then
        return 1
    fi
    return 0
}

# 创建用户
if ! user_exists $worker
then 
    echo "user $worker don't exist, create it"
    useradd -m $worker -p $password
    if ! user_exists $worker
    then 
        echo "user $worker create error"
        exit 1
    fi
fi

su - $worker -c "$runsh $PWD/../config.yaml"
#userdel $worker