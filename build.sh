#!/usr/bin/env bash
#
# Auth:Eric Shi
# Email:shibingli@yeah.net
# QQ:155122504
#
#
DIRNAME=$(cd "$(dirname "$0")"; pwd)

which go > /dev/null 2>&1
if [[ $? -eq 0 ]]; then
    echo ""
    echo "==> Golang is already installed."
else
    echo "==> Install Golang."
    echo ""
    tar -zxvf $DIRNAME/go1.7.3.linux-amd64.tar.gz -C /usr/local
    echo ""
    export PATH=/usr/local/go/bin:$PATH
fi

which go > /dev/null 2>&1
if [[ $? -eq 0 ]]; then
    export GOPATH=$DIRNAME
    git submodule update --init --recursive
    cd $GOPATH/src/apibox.club/apibox/ && go install
    echo ""
    echo "==> Please run the $GOPATH/bin/apibox."
    echo ""
else
    echo "==> Please install Golang."
fi

