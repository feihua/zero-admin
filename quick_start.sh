#!/bin/bash

#停止服务
docker stop sys
docker stop ums
docker stop oms
docker stop pms
docker stop sms
docker stop cms
docker stop search
docker stop admin-api
docker stop front-api
docker stop job
docker stop consumer

#删除容器
docker rm sys
docker rm ums
docker rm oms
docker rm pms
docker rm sms
docker rm cms
docker rm search
docker rm admin-api
docker rm front-api
docker rm job
docker rm consumer

#删除镜像
docker rmi sys-rpc:0.0.1
docker rmi ums-rpc:0.0.1
docker rmi oms-rpc:0.0.1
docker rmi pms-rpc:0.0.1
docker rmi sms-rpc:0.0.1
docker rmi cms-rpc:0.0.1
docker rmi search-rpc:0.0.1
docker rmi admin-api:0.0.1
docker rmi front-api:0.0.1
docker rm job:0.0.1
docker rm consumer:0.0.1


#删除none镜像
docker rmi $(docker images | grep "none" | awk '{print $3}')

#构建服务
docker build -t sys-rpc:0.0.1 -f rpc/sys/Dockerfile .
docker build -t ums-rpc:0.0.1 -f rpc/ums/Dockerfile .
docker build -t oms-rpc:0.0.1 -f rpc/oms/Dockerfile .
docker build -t pms-rpc:0.0.1 -f rpc/pms/Dockerfile .
docker build -t sms-rpc:0.0.1 -f rpc/sms/Dockerfile .
docker build -t cms-rpc:0.0.1 -f rpc/cms/Dockerfile .
docker build -t search-rpc:0.0.1 -f rpc/search/Dockerfile .
docker build -t admin-api:0.0.1 -f api/admin/Dockerfile .
docker build -t front-api:0.0.1 -f api/front/Dockerfile .
docker build -t job:0.0.1 -f job/Dockerfile .
docker build -t consumer:0.0.1 -f consumer/Dockerfile .

#启动服务
docker run -itd --net=host --name=sys sys-rpc:0.0.1
docker run -itd --net=host --name=ums ums-rpc:0.0.1
docker run -itd --net=host --name=oms oms-rpc:0.0.1
docker run -itd --net=host --name=pms pms-rpc:0.0.1
docker run -itd --net=host --name=sms sms-rpc:0.0.1
docker run -itd --net=host --name=cms cms-rpc:0.0.1
docker run -itd --net=host --name=search search-rpc:0.0.1
docker run -itd --net=host --name=admin-api  -v /www/wwwroot/upload:/app/uploads admin-api:0.0.1
docker run -itd --net=host --name=front-api front-api:0.0.1
docker run -itd --net=host --name=job job:0.0.1
docker run -itd --net=host --name=consumer consumer:0.0.1
