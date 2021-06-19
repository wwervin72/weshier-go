#!/bin/bash

pkgName="weshiergo"
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

docker run -p 3333:3333 --restart always -v /data/weshier/log:/build/log -v /data/weshier/assets/public:/build/public -v /data/weshier/conf/config.json:/build/conf/config.json -d --network ${net} --name ${pkgName} weshiergo:v0.0.1
