#!/bin/bash
echo ''
echo 'File to go protobuf'
array=()
if [ $1 ]; then
    if [ -e $1 ]; then
        if [ -d $1 ]; then
            array=`find $1 -name *.proto`
        elif [ -f $1 ]; then
            a=$(expr $1 : '.*\.proto$')
            if [ ${a} != 0 ]; then
                array=($1)
            fi
        fi
    fi
else
    array=`find . -name *.proto`
fi

for path in ${array[*]}; do
    protoc --go_out=../models ${path}
done
echo 'All jobs done'
echo ''