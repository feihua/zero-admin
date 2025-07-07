# 定义变量
ifndef GOPATH
	GOPATH := $(shell go env GOPATH)
endif

GOBIN=$(GOPATH)/bin
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) mod tidy

GOCTL=$(GOBIN)/goctl ## goctl

# 安装goctl代码生成工具
$(shell if [ ! -d $(GOCTL) ]; then \
	$(GOCMD) install github.com/zeromicro/go-zero/tools/goctl@latest; \
fi; \
)

all: deps build ## 默认的构建目标


clean: ## 清理目标
	$(GOCLEAN)
	rm -rf target
 

deps: ## 安装依赖目标
	@export GOPROXY=https://goproxy.cn,direct
	$(GOGET) -v
 
copy_config:
	mkdir -p target/sys-rpc && cp rpc/sys/etc/sys.yaml target/sys-rpc/sys-rpc.yaml
	mkdir -p target/ums-rpc && cp rpc/ums/etc/ums.yaml target/ums-rpc/ums-rpc.yaml
	mkdir -p target/oms-rpc && cp rpc/oms/etc/oms.yaml target/oms-rpc/oms-rpc.yaml
	mkdir -p target/pms-rpc && cp rpc/pms/etc/pms.yaml target/pms-rpc/pms-rpc.yaml
	mkdir -p target/cms-rpc && cp rpc/cms/etc/cms.yaml target/cms-rpc/cms-rpc.yaml
	mkdir -p target/sms-rpc && cp rpc/sms/etc/sms.yaml target/sms-rpc/sms-rpc.yaml
	mkdir -p target/search-rpc && cp rpc/search/etc/search.yaml target/search-rpc/search-rpc.yaml
	mkdir -p target/admin-api && cp api/admin/etc/admin-api.yaml target/admin-api/admin-api.yaml
	mkdir -p target/front-api && cp api/front/etc/front-api.yaml target/front-api/front-api.yaml
	mkdir -p target/job && cp job/etc/job.yaml target/job/job.yaml
	mkdir -p target/consumer && cp consumer/etc/consumer.yaml target/consumer/consumer.yaml

build: copy_config ## 构建目标
	$(GOBUILD) -o target/sys-rpc/sys-rpc -v ./rpc/sys/sys.go
	$(GOBUILD) -o target/ums-rpc/ums-rpc -v ./rpc/ums/ums.go
	$(GOBUILD) -o target/oms-rpc/oms-rpc -v ./rpc/oms/oms.go
	$(GOBUILD) -o target/pms-rpc/pms-rpc -v ./rpc/pms/pms.go
	$(GOBUILD) -o target/cms-rpc/cms-rpc -v ./rpc/cms/cms.go
	$(GOBUILD) -o target/sms-rpc/sms-rpc -v ./rpc/sms/sms.go
	$(GOBUILD) -o target/search-rpc/search-rpc -v ./rpc/search/search.go
	$(GOBUILD) -o target/admin-api/admin-api -v ./api/admin/admin.go
	$(GOBUILD) -o target/front-api/front-api -v ./api/front/front.go
	$(GOBUILD) -o target/job/job -v ./job/job.go
	$(GOBUILD) -o target/consumer/consumer -v ./consumer/consumer.go


start: ## 运行目标
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
	nohup ./target/job/job -f ./target/job/job.yaml  > /dev/null 2>&1 &
	nohup ./target/consumer/consumer -f ./target/consumer/consumer.yaml  > /dev/null 2>&1 &


stop: ## 停止目标
	-pkill -f admin-api
	-pkill -f front-api
	-pkill -f consumer
	-pkill -f job
	-pkill -f sys-rpc
	-pkill -f sms-rpc
	-pkill -f cms-rpc
	-pkill -f ums-rpc
	-pkill -f oms-rpc
	-pkill -f pms-rpc
	-pkill -f search-rpc
	@for i in 5 4 3 2 1; do\
      echo -n "stop $$i";\
      sleep 1; \
      echo " "; \
    done


restart: stop start ## 重启项目

test: build restart ## 快速测试


.DEFAULT_GOAL := all ## 默认构建目标是


format: ## 格式化代码
	$(GOCTL) api format --dir api/admin/doc/api
	$(GOCTL) api format --dir api/front/doc/api

gen:	## 生成所有模块代码
	$(GOCTL) api go -api ./api/admin/doc/api/admin.api -dir ./api/admin/
	# 生成front-api代码
	$(GOCTL) api go -api ./api/front/doc/api/front.api -dir ./api/front/ --style go_zero
	# 生成mq消费者代码
	$(GOCTL) api go -api ./consumer/consumer.api -dir ./consumer --style go_zero
	# 生成定时任务代码
	$(GOCTL) api go -api ./job/job.api -dir ./job --style go_zero
	# 生成sys-rpc代码
	$(GOCTL) rpc protoc rpc/sys/sys.proto --go_out=./rpc/sys/ --go-grpc_out=./rpc/sys/ --zrpc_out=./rpc/sys/ -m
	# 生成ums-rpc代码
	$(GOCTL) rpc protoc rpc/ums/ums.proto --go_out=./rpc/ums/ --go-grpc_out=./rpc/ums/ --zrpc_out=./rpc/ums/ -m
	# pms-rpc代码
	$(GOCTL) rpc protoc rpc/pms/pms.proto --go_out=./rpc/pms/ --go-grpc_out=./rpc/pms/ --zrpc_out=./rpc/pms/ -m
	# 生成oms-rpc代码
	$(GOCTL) rpc protoc rpc/oms/oms.proto --go_out=./rpc/oms/ --go-grpc_out=./rpc/oms/ --zrpc_out=./rpc/oms/ -m --style go_zero
	# 生成sms-rpc代码
	$(GOCTL) rpc protoc rpc/sms/sms.proto --go_out=./rpc/sms/ --go-grpc_out=./rpc/sms/ --zrpc_out=./rpc/sms/ -m --style go_zero
	# 生成cmsrpc代码
	$(GOCTL) rpc protoc rpc/cms/cms.proto --go_out=./rpc/cms/ --go-grpc_out=./rpc/cms/ --zrpc_out=./rpc/cms/ -m
	# 生成search代码
	$(GOCTL) rpc protoc rpc/search/search.proto --go_out=./rpc/search/ --go-grpc_out=./rpc/search/ --zrpc_out=./rpc/search/ --style go_zero


model: ## 生成model代码
	go run rpc/cms/gen/generator.go
	go run rpc/oms/gen/generator.go
	go run rpc/pms/gen/generator.go
	go run rpc/sms/gen/generator.go
	go run rpc/sys/gen/generator.go
	go run rpc/ums/gen/generator.go


image: ## 构建docker镜像
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

run: ## 启动docker容器
	docker run -itd --net=host --name=sys sys-rpc:0.0.1; \
    docker run -itd --net=host --name=ums ums-rpc:0.0.1; \
    docker run -itd --net=host --name=oms oms-rpc:0.0.1; \
    docker run -itd --net=host --name=pms pms-rpc:0.0.1; \
    docker run -itd --net=host --name=sms sms-rpc:0.0.1; \
    docker run -itd --net=host --name=cms cms-rpc:0.0.1; \
    docker run -itd --net=host --name=search search-rpc:0.0.1; \
    docker run -itd --net=host --name=admin-api admin-api:0.0.1; \
    docker run -itd --net=host --name=front-api front-api:0.0.1; \
    docker run -itd --net=host --name=job job:0.0.1 \
    docker run -itd --net=host --name=consumer consumer:0.0.1 \

kubectl: ## 部署k8s容器
	kubectl apply -f script/account/serviceaccount.yaml; \
    kubectl apply -f script/configmap/register.yaml; \
    kubectl apply -f script/sys-rpc.yaml; \
    kubectl apply -f script/ums-rpc.yaml; \
    kubectl apply -f script/sms-rpc.yaml; \
    kubectl apply -f script/pms-rpc.yaml; \
    kubectl apply -f script/oms-rpc.yaml; \
    kubectl apply -f script/cms-rpc.yaml; \
    kubectl apply -f script/search-rpc.yaml; \
    kubectl apply -f script/admin-api.yaml; \
    kubectl apply -f script/front-api.yaml; \
    kubectl apply -f script/job.yaml; \
    kubectl apply -f script/consumer.yaml; \

help: ## show help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% 0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
