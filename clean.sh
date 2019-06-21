#!/bin/bash

docker rm -f $(docker ps -aq)
images=( realestate-peer orderer regulator-peer trader-ca regulator-ca realestate-ca trader-peer )
for i in "${images[@]}"
do
	echo Removing image : $i
  docker rmi -f $i
done

#docker rmi -f $(docker images | grep none)
images=( dev-trader-peer dev-regulator-peer dev-realestate-peer dev-shop-peer)
for i in "${images[@]}"
do
	echo Removing image : $i
  docker rmi -f $(docker images | grep $i )
done

#!/bin/sh
rm -rf ./web/realestate-peer
rm -rf ./web/regulator-peer
rm -rf ./web/trader-peer
