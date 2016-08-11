#!/bin/bash
echo ''
echo 'File to go protobuf'
echo ''
protoGoFile(){
    protoc --go_out=../models/ $1
}

rangeDir(){
    for path in `ls $1 |grep .proto`; do
        if [ -d ${path} ]; then
            echo 'TODO: This is a dir'
            #rangeDir ${path}
        else
            protoGoFile ${path}
        fi
    done
}

if [ $1 ]; then
    if [ -e $1 ]; then
        if [ -d $1 ]; then
            echo 'TODO: This is a dir'
        else
            protoGoFile $1
        fi
    fi
else
    rangeDir `pwd`
fi

echo ''
echo 'All jobs done'
echo ''