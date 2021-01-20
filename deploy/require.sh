#!/bin/bash
# please run with root or sudo

# 判断程序是否存在
program_exists() {
    local ret='0'
    command -v $1 >/dev/null 2>&1 || { local ret='1'; }

    # fail on non-zero return value
    if [ "$ret" -ne 0 ]; then
        return 1
    fi
    return 0
}

if ! program_exists go
then
    wget -O go.tar.gz https://dl.google.com/go/go1.15.7.linux-amd64.tar.gz
    tar -C /usr/local -xzf go.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    if ! program_exists go
    then
        echo "go install error"
        exit 1
    fi
else
    echo "gogogo"
fi