# 定义变量
ifndef GOPATH
	GOPATH := $(shell go env GOPATH)
endif

GOBIN=$(GOPATH)/bin
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

# goctl
GOCTL=$(GOBIN)/goctl

# 生成admin-api代码
admin:
	$(GOCTL) api go -api ./api/doc/api/admin.api -dir ./api/

# 生成front-api代码
front:
	/$(GOCTL) api go -api ./front-api/doc/api/front.api -dir ./front-api/

# 生成sys-rpc代码
sys:
	$(GOCTL) rpc protoc rpc/sys/sys.proto --go_out=./rpc/sys/ --go-grpc_out=./rpc/sys/ --zrpc_out=./rpc/sys/ -m

# 生成ums-rpc代码
ums:
	$(GOCTL) rpc protoc rpc/ums/ums.proto --go_out=./rpc/ums/ --go-grpc_out=./rpc/ums/ --zrpc_out=./rpc/ums/ -m

# pms-rpc代码
pms:
	$(GOCTL) rpc protoc rpc/pms/pms.proto --go_out=./rpc/pms/ --go-grpc_out=./rpc/pms/ --zrpc_out=./rpc/pms/ -m

# 生成oms-rpc代码
oms:
	$(GOCTL) rpc protoc rpc/oms/oms.proto --go_out=./rpc/oms/ --go-grpc_out=./rpc/oms/ --zrpc_out=./rpc/oms/ -m

# 生成sms-rpc代码
sms:
	$(GOCTL) rpc protoc rpc/sms/sms.proto --go_out=./rpc/sms/ --go-grpc_out=./rpc/sms/ --zrpc_out=./rpc/sms/ -m

# 生成cmsrpc代码
cms:
	$(GOCTL) rpc protoc rpc/cms/cms.proto --go_out=./rpc/cms/ --go-grpc_out=./rpc/cms/ --zrpc_out=./rpc/cms/ -m

# 生成所有模块代码
gen:
	$(GOCTL) api go -api ./api/doc/api/admin.api -dir ./api/
	# 生成front-api代码
	/$(GOCTL) api go -api ./front-api/doc/api/front.api -dir ./front-api/
	# 生成sys-rpc代码
	$(GOCTL) rpc protoc rpc/sys/sys.proto --go_out=./rpc/sys/ --go-grpc_out=./rpc/sys/ --zrpc_out=./rpc/sys/ -m
	# 生成ums-rpc代码
	$(GOCTL) rpc protoc rpc/ums/ums.proto --go_out=./rpc/ums/ --go-grpc_out=./rpc/ums/ --zrpc_out=./rpc/ums/ -m
	# pms-rpc代码
	$(GOCTL) rpc protoc rpc/pms/pms.proto --go_out=./rpc/pms/ --go-grpc_out=./rpc/pms/ --zrpc_out=./rpc/pms/ -m
	# 生成oms-rpc代码
	$(GOCTL) rpc protoc rpc/oms/oms.proto --go_out=./rpc/oms/ --go-grpc_out=./rpc/oms/ --zrpc_out=./rpc/oms/ -m
	# 生成sms-rpc代码
	$(GOCTL) rpc protoc rpc/sms/sms.proto --go_out=./rpc/sms/ --go-grpc_out=./rpc/sms/ --zrpc_out=./rpc/sms/ -m
	# 生成cmsrpc代码
	$(GOCTL) rpc protoc rpc/cms/cms.proto --go_out=./rpc/cms/ --go-grpc_out=./rpc/cms/ --zrpc_out=./rpc/cms/ -m