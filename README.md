# Cài đặt mạng Revex trên local
Clone mã nguồn tại https://github.com/tuandht97/revex.git

```
cd revex 
# Xóa toàn bộ các container và image docker trên máy
./clean.sh 
# Khởi tạo lại các image docker
./docker-images.sh 
# Khởi tạo các container được cấu hình trong file docker-composer.yaml
docker-compose up -d 
# Truy cập vào command line của cli container
docker exec -it cli bash 
```

*Trong giao diện command line của cli* 

```
# Sao chép file script.sh
cp scripts/script.sh .
# Cấp quyền sử dụng file script.sh
chmod +x script.sh 
# Chạy file script.sh để cài đặt chaincode lên các container trong mạng
./script.sh 
```

## Chạy web-service
Clone web-service tại https://github.com/tuandht97/revex-service.git và làm theo hướng dẫn

# Cài đặt mạng revex trên google cloud

Hệ thống được cấu hình trên 4 máy
Clone mã nguồn trên cả 4 máy tại https://github.com/tuandht97/revex.git

**Xóa bỏ các service và container**

Sử dụng các lệnh sau trên từng máy để xóa các service và container docker đang mở
```
docker service rm $(docker service ls -q)
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
```
**Tại máy tính 1 (chain-1)**
```
cd revex
# Tạo các service docker được cấu hình trong file yaml
docker stack deploy --compose-file docker-compose_dev_1.yaml revex
```

**Tại máy tính 2 (chain-2)**
```
cd revex
# Tạo các service docker được cấu hình trong file yaml
docker stack deploy --compose-file docker-compose_dev_2.yaml revex
```

**Tại máy tính 3 (chain-3)**
```
cd revex
# Tạo các service docker được cấu hình trong file yaml
docker stack deploy --compose-file docker-compose_dev_3.yaml revex
```

**Tại máy tính 4 (chain-4)**
```
cd revex
# Tạo các service docker được cấu hình trong file yaml
docker stack deploy --compose-file docker-compose_dev_4.yaml revex
```

**Khởi tạo mạng với CLI**

Cli container được cấu hình chạy trên máy 3 (chain-3)

```
# Truy cập vào command line của cli container
docker exec -it <id của container tìm đc> bash

# Sao chép file script.sh
cp scripts/script.sh .
# Cấp quyền sử dụng file script.sh
chmod +x script.sh 
# Chạy file script.sh để cài đặt chaincode lên các container trong mạng
./script.sh 
```
