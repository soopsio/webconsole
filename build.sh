#!/usr/bin/env sh
#
# Auth:Eric Shi
# Email:shibingli@yeah.net
# QQ:155122504
#
#
DIRNAME=$(cd "$(dirname "$0")"; pwd)

which go > /dev/null 2>&1
if [ $? == 0 ]; then
    echo ""
    echo "==> Golang is already installed."
else
    echo "==> Download Golang."
    echo ""
    curl -L "https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz" | tar -zx -C /usr/local
    echo ""
    PATH=/usr/local/go/bin:$PATH
    export $PATH
fi

which go > /dev/null 2>&1
if [ $? == 0 ]; then
    GOPATH=$DIRNAME
    cd $GOPATH/src/apibox.club/apibox/ && go install
    echo ""
    echo "==> Please run the $GOPATH/bin/apibox."
    echo ""
else
    echo "==> Please install Golang."
fi

