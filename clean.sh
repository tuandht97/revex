#!/bin/bash

docker stop -f $(docker ps -aq)
docker rm -f $(docker ps -aq)
images=( bds-peer orderer ccq-peer trader-ca ccq-ca bds-ca trader-peer )
for i in "${images[@]}"
do
	echo Removing image : $i
  docker rmi -f $i
done

#docker rmi -f $(docker images | grep none)
images=( dev-trader-peer dev-ccq-peer dev-bds-peer dev-shop-peer)
for i in "${images[@]}"
do
	echo Removing image : $i
  docker rmi -f $(docker images | grep $i )
done

