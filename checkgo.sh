#!/bin/bash
echo ''
echo 'Format code and import'
array=()
for path in `go list ./... | grep -v vendor`; do
    newPath=${GOPATH}/src/${path}
    checkPath=${path}/vendor/
    if [ -d ${newPath}/vendor ]; then
        if [ ${#array[@]} = 0 ]; then
            array[${#array[@]}] = ${checkPath}
        else
            for item in array; do
            if [ ${item} = ${checkPath} ]; then
                break
            fi
            done
            array[${#array[@]}] = ${checkPath}
        fi
        continue
    fi
    gofmt -w ${newPath}
    goimports -w ${newPath}
done
for item in ${array[*]}; do
    change=`echo ${item} | sed 's#\/#\\\/#g'`
    sed -i 's/'"${change}"'//g' `find ./ -type f |grep -v vendor/ |grep -v .git |grep -v Godeps/`
done

echo ''
echo 'Go vet `go list ./... |grep -v vendor/`'
go vet `go list ./... |grep -v vendor/`

echo ''
echo 'Golint ./... |grep -v vendor/'
golint ./... |grep -v vendor/

echo ''
echo 'Check go all jobs done!!!'
echo ''