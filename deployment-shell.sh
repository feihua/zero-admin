#!/bin/bash

#停止服务
docker stop sys-rpc
docker stop ums-rpc
docker stop oms-rpc
docker stop pms-rpc
docker stop sms-rpc
docker stop cms-rpc
docker stop admin-api
docker stop front-api

#删除容器
docker rm sys-rpc
docker rm ums-rpc
docker rm oms-rpc
docker rm pms-rpc
docker rm sms-rpc
docker rm cms-rpc
docker rm admin-api
docker rm front-api

#删除镜像
docker rmi sys-rpc:0.0.1
docker rmi ums-rpc:0.0.1
docker rmi oms-rpc:0.0.1
docker rmi pms-rpc:0.0.1
docker rmi sms-rpc:0.0.1
docker rmi cms-rpc:0.0.1
docker rmi admin-api:0.0.1
docker rmi front-api:0.0.1

#删除none镜像
docker rmi $(docker images | grep "none" | awk '{print $3}')

#构建服务
docker build -t sys-rpc:0.0.1 -f rpc/sys/Dockerfile .
docker build -t ums-rpc:0.0.1 -f rpc/ums/Dockerfile .
docker build -t oms-rpc:0.0.1 -f rpc/oms/Dockerfile .
docker build -t pms-rpc:0.0.1 -f rpc/pms/Dockerfile .
docker build -t sms-rpc:0.0.1 -f rpc/sms/Dockerfile .
docker build -t cms-rpc:0.0.1 -f rpc/cms/Dockerfile .
docker build -t admin-api:0.0.1 -f api/Dockerfile .
docker build -t front-api:0.0.1 -f front-api/Dockerfile .

#启动服务
docker run -itd --net=host --name=sys sys-rpc:0.0.1
docker run -itd --net=host --name=ums ums-rpc:0.0.1
docker run -itd --net=host --name=oms oms-rpc:0.0.1
docker run -itd --net=host --name=pms pms-rpc:0.0.1
docker run -itd --net=host --name=sms sms-rpc:0.0.1
docker run -itd --net=host --name=cms cms-rpc:0.0.1
docker run -itd --net=host --name=admin-api admin-api:0.0.1
docker run -itd --net=host --name=front-api front-api:0.0.1
