#!/bin/bash
# Build gosdk project and run the testing
# Author: songshuangyang@baidu.com
# Time:   2017.8

CURRENT=$(pwd)
if [ -a $CURRENT/pkg ];then
    rm -r $CURRENT/pkg
fi
GOBIN=`which go`
if [ $? -ne 0 ];then
    echo "can not find the executable go"
    exit 1
fi
echo "found the executable go: $GOBIN"

export GOPATH=$CURRENT
cd $CURRENT

go install baidubce/auth
echo "install the baidubce/auth package"
go install baidubce/bce
echo "install the baidubce/bce package"
go install baidubce/http
echo "install the baidubce/http package"
go install baidubce/util
echo "install the baidubce/util package"
go install baidubce/services/bos
echo "install the baidubce/services/bos package"
go install baidubce/services/sts
echo "install the baidubce/services/sts package"

echo "Build gosdk success"

if [ -z $1 ];then
    exit 0
fi
if [ $1 == "test" ];then

    echo "starting test BOS service..."
    go test -v baidubce/services/bos

    echo "starting test STS service..."
    go test -v baidubce/services/sts
fi
exit 0

