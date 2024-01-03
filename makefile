# 定义变量
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) mod tidy

# 默认的构建目标
all: clean deps build run

# 清理目标
clean:
	$(GOCLEAN)
	rm -rf target
 
# 安装依赖目标
deps:
	@export GOPROXY=https://goproxy.cn,direct
	$(GOGET) -v
 
# 构建目标
build:
	$(GOBUILD) -o target/admin-api -v ./api/admin.go
	$(GOBUILD) -o target/front-api -v ./front-api/front.go
	$(GOBUILD) -o target/sys-rpc -v ./rpc/sys/sys.go
	$(GOBUILD) -o target/ums-rpc -v ./rpc/ums/ums.go
	$(GOBUILD) -o target/oms-rpc -v ./rpc/oms/oms.go
	$(GOBUILD) -o target/pms-rpc -v ./rpc/pms/pms.go
	$(GOBUILD) -o target/cms-rpc -v ./rpc/cms/cms.go
	$(GOBUILD) -o target/sms-rpc -v ./rpc/sms/sms.go

# 运行目标
run:
	nohup ./target/admin-api -f api/etc/admin-api.yaml > /dev/null 2>&1 &
	nohup ./target/front-api -f front-api/etc/front-api.yaml  > /dev/null 2>&1 &
	nohup ./target/sys-rpc -f rpc/sys/etc/sys.yaml  > /dev/null 2>&1 &
	nohup ./target/ums-rpc -f rpc/ums/etc/ums.yaml  > /dev/null 2>&1 &
	nohup ./target/oms-rpc -f rpc/oms/etc/oms.yaml  > /dev/null 2>&1 &
	nohup ./target/pms-rpc -f rpc/pms/etc/pms.yaml  > /dev/null 2>&1 &
	nohup ./target/cms-rpc -f rpc/cms/etc/cms.yaml  > /dev/null 2>&1 &
	nohup ./target/sms-rpc -f rpc/sms/etc/sms.yaml  > /dev/null 2>&1 &

# 停止目标
stop:
	pkill -f admin-api
	pkill -f front-api
	pkill -f sys-rpc
	pkill -f sms-rpc
	pkill -f cms-rpc
	pkill -f ums-rpc
	pkill -f oms-rpc
	pkill -f pms-rpc 

# 重启
restart: stop run

# 默认构建目标是 "all"
.DEFAULT_GOAL := all
