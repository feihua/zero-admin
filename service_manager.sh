#!/bin/bash

GOCMD="go"
GOBUILD="$GOCMD build"
GOCLEAN="$GOCMD clean"
GOGET="$GOCMD mod tidy"

export GOPROXY=https://goproxy.cn,direct
go mod tidy

# 生成配置文件
config(){
  mkdir -p target/sys-rpc && cp -n rpc/sys/etc/sys.yaml target/sys-rpc/sys-rpc.yaml
	mkdir -p target/ums-rpc && cp -n rpc/ums/etc/ums.yaml target/ums-rpc/ums-rpc.yaml
	mkdir -p target/oms-rpc && cp -n rpc/oms/etc/oms.yaml target/oms-rpc/oms-rpc.yaml
	mkdir -p target/pms-rpc && cp -n rpc/pms/etc/pms.yaml target/pms-rpc/pms-rpc.yaml
	mkdir -p target/cms-rpc && cp -n rpc/cms/etc/cms.yaml target/cms-rpc/cms-rpc.yaml
	mkdir -p target/sms-rpc && cp -n rpc/sms/etc/sms.yaml target/sms-rpc/sms-rpc.yaml
	mkdir -p target/search-rpc && cp -n rpc/search/etc/search.yaml target/search-rpc/search-rpc.yaml
	mkdir -p target/admin-api && cp -n api/admin/etc/admin-api.yaml target/admin-api/admin-api.yaml
	mkdir -p target/front-api && cp -n api/front/etc/front-api.yaml target/front-api/front-api.yaml
	mkdir -p target/job && cp -n job/etc/job-api.yaml target/job/job-api.yaml
	mkdir -p target/consumer && cp -n consumer/etc/consumer-api.yaml target/consumer/consumer-api.yaml

}
# 编译
build(){
  	${GOBUILD} -o target/sys-rpc/sys-rpc -v ./rpc/sys/sys.go
  	${GOBUILD} -o target/ums-rpc/ums-rpc -v ./rpc/ums/ums.go
  	${GOBUILD} -o target/oms-rpc/oms-rpc -v ./rpc/oms/oms.go
  	${GOBUILD} -o target/pms-rpc/pms-rpc -v ./rpc/pms/pms.go
  	${GOBUILD} -o target/cms-rpc/cms-rpc -v ./rpc/cms/cms.go
  	${GOBUILD} -o target/sms-rpc/sms-rpc -v ./rpc/sms/sms.go
  	${GOBUILD} -o target/search-rpc/search-rpc -v ./rpc/search/search.go
  	${GOBUILD} -o target/admin-api/admin-api -v ./api/admin/admin.go
  	${GOBUILD} -o target/front-api/front-api -v ./api/front/front.go
  	${GOBUILD} -o target/job/job -v ./job/job.go
  	${GOBUILD} -o target/consumer/consumer -v ./consumer/consumer.go
}

# 启动
start(){
  	nohup ./target/sys-rpc/sys-rpc -f ./target/sys-rpc/sys-rpc.yaml  > /dev/null 2>&1 &
  	nohup ./target/ums-rpc/ums-rpc -f ./target/ums-rpc/ums-rpc.yaml  > /dev/null 2>&1 &
  	nohup ./target/oms-rpc/oms-rpc -f ./target/oms-rpc/oms-rpc.yaml  > /dev/null 2>&1 &
  	nohup ./target/pms-rpc/pms-rpc -f ./target/pms-rpc/pms-rpc.yaml  > /dev/null 2>&1 &
  	nohup ./target/cms-rpc/cms-rpc -f ./target/cms-rpc/cms-rpc.yaml  > /dev/null 2>&1 &
  	nohup ./target/sms-rpc/sms-rpc -f ./target/sms-rpc/sms-rpc.yaml  > /dev/null 2>&1 &
  	nohup ./target/search-rpc/search-rpc -f ./target/search-rpc/search-rpc.yaml  > /dev/null 2>&1 &
  	nohup ./target/admin-api/admin-api -f ./target/admin-api/admin-api.yaml > /dev/null 2>&1 &
  	nohup ./target/front-api/front-api -f ./target/front-api/front-api.yaml  > /dev/null 2>&1 &
  	nohup ./target/web-api/web-api -f ./target/web-api/web-api.yaml  > /dev/null 2>&1 &
  	nohup ./target/job/job -f ./target/job/job-api.yaml  > /dev/null 2>&1 &
  	nohup ./target/consumer/consumer -f ./target/consumer/consumer-api.yaml  > /dev/null 2>&1 &

}

# 停止
stop_time(){
  for i in 5 4 3 2 1; do
      echo -n "stop $i";
      sleep 1;
      echo " ";
  done
}
# 停止服务
stop(){
	pkill -f admin-api
	pkill -f front-api
	pkill -f consumer
	pkill -f job
	pkill -f sys-rpc
	pkill -f sms-rpc
	pkill -f cms-rpc
	pkill -f ums-rpc
	pkill -f oms-rpc
	pkill -f pms-rpc
	pkill -f search-rpc
	stop_time
}

# 系统服务
sys_service(){
  mkdir -p target/sys-rpc && cp -n rpc/sys/etc/sys.yaml target/sys-rpc/sys-rpc.yaml
  go build -o target/sys-rpc/sys-rpc -v ./rpc/sys/sys.go
  pkill -f sys-rpc
  stop_time
  nohup ./target/sys-rpc/sys-rpc -f ./target/sys-rpc/sys-rpc.yaml  > /dev/null 2>&1 &
}

# 用户服务
ums_service(){
  mkdir -p target/ums-rpc && cp -n rpc/ums/etc/ums.yaml target/ums-rpc/ums-rpc.yaml
  ${GOBUILD} -o target/ums-rpc/ums-rpc -v ./rpc/ums/ums.go
  pkill -f ums-rpc
  stop_time
  nohup ./target/ums-rpc/ums-rpc -f ./target/ums-rpc/ums-rpc.yaml  > /dev/null 2>&1 &
}

# 订单服务
oms_service(){
  mkdir -p target/oms-rpc && cp -n rpc/oms/etc/oms.yaml target/oms-rpc/oms-rpc.yaml
  ${GOBUILD} -o target/oms-rpc/oms-rpc -v ./rpc/oms/oms.go
  pkill -f oms-rpc
  stop_time
  nohup ./target/oms-rpc/oms-rpc -f ./target/oms-rpc/oms-rpc.yaml  > /dev/null 2>&1 &
}

# 商品服务
pms_service(){
  mkdir -p target/pms-rpc && cp -n rpc/pms/etc/pms.yaml target/pms-rpc/pms-rpc.yaml
  ${GOBUILD} -o target/pms-rpc/pms-rpc -v ./rpc/pms/pms.go
  pkill -f pms-rpc
  stop_time
  nohup ./target/pms-rpc/pms-rpc -f ./target/pms-rpc/pms-rpc.yaml  > /dev/null 2>&1 &
}

# 内容服务
cms_service(){
  mkdir -p target/cms-rpc && cp -n rpc/cms/etc/cms.yaml target/cms-rpc/cms-rpc.yaml
  ${GOBUILD} -o target/cms-rpc/cms-rpc -v ./rpc/cms/cms.go
  pkill -f cms-rpc
  stop_time
  nohup ./target/cms-rpc/cms-rpc -f ./target/cms-rpc/cms-rpc.yaml  > /dev/null 2>&1 &
}

# 营销服务
sms_service(){
  mkdir -p target/sms-rpc && cp -n rpc/sms/etc/sms.yaml target/sms-rpc/sms-rpc.yaml
  ${GOBUILD} -o target/sms-rpc/sms-rpc -v ./rpc/sms/sms.go
  pkill -f sms-rpc
  stop_time
  nohup ./target/sms-rpc/sms-rpc -f ./target/sms-rpc/sms-rpc.yaml  > /dev/null 2>&1 &
}

# 搜索服务
search_service(){
  mkdir -p target/search-rpc && cp -n rpc/search/etc/search.yaml target/search-rpc/search-rpc.yaml
  ${GOBUILD} -o target/search-rpc/search-rpc -v ./rpc/search/search.go
  pkill -f search-rpc
  stop_time
  nohup ./target/search-rpc/search-rpc -f ./target/search-rpc/search-rpc.yaml  > /dev/null 2>&1 &
}

# 后台服务
admin_service(){
  mkdir -p target/admin-api && cp -n api/admin/etc/admin-api.yaml target/admin-api/admin-api.yaml
  ${GOBUILD} -o target/admin-api/admin-api -v ./api/admin/admin.go
  pkill -f admin-api
  stop_time
  nohup ./target/admin-api/admin-api -f ./target/admin-api/admin-api.yaml > /dev/null 2>&1 &
}

# 前台服务
front_service(){
  mkdir -p target/front-api && cp -n api/front/etc/front-api.yaml target/front-api/front-api.yaml
  ${GOBUILD} -o target/front-api/front-api -v ./api/front/front.go
  pkill -f front-api
  stop_time
  nohup ./target/front-api/front-api -f ./target/front-api/front-api.yaml  > /dev/null 2>&1 &
}

# 任务服务
job_service(){
  mkdir -p target/job && cp -n job/etc/job-api.yaml target/job/job-api.yaml
  ${GOBUILD} -o target/job/job -v ./job/job.go
  pkill -f job
  stop_time
  nohup ./target/job/job -f ./target/job/job-api.yaml  > /dev/null 2>&1 &
}

# 消费者服务
consumer_service(){
  mkdir -p target/consumer && cp -n consumer/etc/consumer-api.yaml target/consumer/consumer-api.yaml
  ${GOBUILD} -o target/consumer/consumer -v ./consumer/consumer.go
  pkill -f consumer
  stop_time
  nohup ./target/consumer/consumer -f ./target/consumer/consumer-api.yaml  > /dev/null 2>&1 &
}

show_help(){
   echo "没匹配中命令，请详细参考case的参数"
}

# 主程序逻辑
case "$1" in
    "build")
        config
        build
        ;;
    "start")
        start
        ;;
    "stop")
        stop
        ;;
    "restart")
        stop
        start
        ;;
    "sys")
        sys_service
        ;;
    "ums")
        ums_service
        ;;
    "oms")
        oms_service
        ;;
    "pms")
        oms_service
        ;;
    "cms")
        cms_service
        ;;
    "sms")
        sms_service
        ;;
    "search")
        search_service
        ;;
    "admin")
        admin_service
        ;;
      
    "front")
        front_service
        ;;
    "job")
        job_service
        ;;
    "consumer")
        consumer_service
        ;;
    *)
        show_help
        ;;
esac
