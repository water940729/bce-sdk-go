#!/bin/bash
# Build gosdk project for publish
# Author: songshuangyang@baidu.com
# Time:   2017.8
#         2018.1 Modified

WORKROOT=$(pwd)
cd $WORKROOT
OUTPUT=$WORKROOT/output
if [ -d $OUTPUT ]; then
    rm -rf $OUTPUT
fi

DEST=$OUTPUT/github.com/baidubce/bce-sdk-go
mkdir -p $DEST
cp -Rf ./auth $DEST
cp -Rf ./bce $DEST
cp -Rf ./http $DEST
cp -Rf ./services $DEST
cp -Rf ./util $DEST
cp -Rf ./doc $DEST
echo "Build gosdk success"

