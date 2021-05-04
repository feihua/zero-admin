#!/bin/bash

#停止服务
docker stop sys
docker stop ums
docker stop oms
docker stop pms
docker stop sms
docker stop api

#删除容器
docker rm sys
docker rm ums
docker rm oms
docker rm pms
docker rm sms
docker rm api

#删除镜像
docker rmi sys:v1
docker rmi ums:v1
docker rmi oms:v1
docker rmi pms:v1
docker rmi sms:v1
docker rmi api:v1

#删除none镜像
docker rmi $(docker images | grep "none" | awk '{print $3}')

#构建服务
docker build -t sys:v1 -f rpc/sys/Dockerfile .
docker build -t ums:v1 -f rpc/ums/Dockerfile .
docker build -t oms:v1 -f rpc/oms/Dockerfile .
docker build -t pms:v1 -f rpc/pms/Dockerfile .
docker build -t sms:v1 -f rpc/sms/Dockerfile .
docker build -t api:v1 -f api/Dockerfile .

#启动服务
docker run -itd --net=host --name=sys sys:v1
docker run -itd --net=host --name=ums ums:v1
docker run -itd --net=host --name=oms oms:v1
docker run -itd --net=host --name=pms pms:v1
docker run -itd --net=host --name=sms sms:v1
docker run -itd --net=host --name=api api:v1
