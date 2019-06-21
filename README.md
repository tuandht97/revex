# Install Blockchain network local

rename project to 'realestate'  
```
cd realestate  
./clean.sh  
./docker-images.sh  
docker-compose up -d  
docker exec -it cli bash  

```

*In cli containner*  
```
cp scripts/script.sh .
chmod +x script.sh  
./script.sh 
```



# Start service-api

```
cd service-api  
./clean.sh  
node index.js 
```

# Install Blockchain network in google cloud
**TURN OFF OLD SERVICE**
```
docker service rm $(docker service ls -q)
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
```
**Dev 1**
```
cd realestate
git pull origin master
docker stack deploy --compose-file docker-compose_dev_1.yaml trit
```

**Dev 2**
```
cd realestate
git pull origin master
docker stack deploy --compose-file docker-compose_dev_2.yaml trit
```

**Dev 3**
```
cd realestate
git pull origin master
docker stack deploy --compose-file docker-compose_dev_3.yaml trit
```

**Dev 4**
```
cd realestate
git pull origin master
docker stack deploy --compose-file docker-compose_dev_4.yaml trit
```

**Bootrap network by CLI**
```
Chạy câu lệnh 'docker ps' trên 4 máy dev_1, dev_2, dev_3, dev_4
Tìm máy đang chạy container của fabric tool

dùng câu lệnh docker exec -it <id của container tìm đc> bash
cp scripts/script.sh .
chmod +x script.sh  
./script.sh 
```




