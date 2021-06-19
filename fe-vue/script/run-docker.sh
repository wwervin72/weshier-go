#!/bin/bash

pkgName="weshier-fe"
net=$1

if [ -z "${net}" ];then
	net="bridge"
fi

runningContainerId=$(docker container inspect $pkgName -f '{{.Id}}')
if [ ! -z "$runningContainerId" ]; then
    echo Found existing container $runningContainerId
    echo Stopping
    docker container stop $runningContainerId
    echo Removing
    docker container rm $runningContainerId
fi

networkId=$(docker network inspect $net -f '{{.Id}}')
if [ -z "$networkId" ]; then
	docker network create ${net}
fi

echo "pkgName: ${pkgName}"
echo "network: ${net}"

docker run -p 80:80 --restart always -d -v /data/weshier/ngLog/:/var/log/nginx/ --network ${net} --name ${pkgName} weshierfe:v0.0.1
