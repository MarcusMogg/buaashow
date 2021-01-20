#!/bin/bash
cfgfile=$1
echo "config file is $cfgfile"
cd ~;

# 判断进程是否存在
pid_exists() {
    local ret='0'
    ps -p $1 >/dev/null 2>&1 || { local ret='1'; }

    if [ "$ret" -ne 0 ]; then
        return 1
    fi
    return 0
}

# 更新代码
if [ -d buaashow ]
then
    cd buaashow;
    git pull;
else
    git clone https://github.com/MarcusMogg/buaashow;
    cd buaashow;
fi

# 构建代码
if [ ! -d build ]
then
    mkdir build;
fi
cd build;
export PATH=$PATH:/usr/local/go/bin  # for go
export GIN_MODE=release; # for gin
export GO111MODULE=on; # for go mod
export GOPROXY=https://goproxy.cn;
go build -o buaashow ../main.go;

# 拷贝配置文件
cp $cfgfile ./config.yaml
# 删除以前的进程
pid_running=$(cat cur)
echo "last time run pid $pid_running"

if pid_exists $pid_running
then
    echo "kill last time running"
    kill -9 $pid_running
fi
# 重新运行
nohup ./buaashow > xx.log 2>&1 & echo $! > cur