#!/bin/bash
# Build gosdk project for publish
# Author: songshuangyang@baidu.com
# Time:   2017.8

WORKROOT=$(pwd)
cd $WORKROOT
OUTPUT=$WORKROOT/output
if [ -d $OUTPUT ]; then
    rm -rf $OUTPUT
fi

mkdir $OUTPUT
cp -rf ./src $OUTPUT
echo "Build gosdk success"

