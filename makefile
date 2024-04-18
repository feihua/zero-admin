# 定义变量
ifndef GOPATH
	GOPATH := $(shell go env GOPATH)
endif

GOBIN=$(GOPATH)/bin
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) mod tidy

MYSQL_INFO=root:r-wz9wop62956dh5k9ed@tcp(110.41.179.89:30395)/gozero


all: deps build ## 默认的构建目标


clean: ## 清理目标
	$(GOCLEAN)
	rm -rf target
 

deps: ## 安装依赖目标
	@export GOPROXY=https://goproxy.cn,direct
	$(GOGET) -v
 

build: ## 构建目标
	copy rpc\sys\etc\sys.yaml target\sys.yaml
	copy rpc\ums\etc\ums.yaml target\ums.yaml
	copy rpc\oms\etc\oms.yaml target\oms.yaml
	copy rpc\pms\etc\pms.yaml target\pms.yaml
	copy rpc\cms\etc\cms.yaml target\cms.yaml
	copy rpc\sms\etc\sms.yaml target\sms.yaml
	copy api\admin\etc\admin-api.yaml target\admin-api.yaml
	copy api\web\etc\web-api.yaml target\web-api.yaml
	copy api\front\etc\front-api.yaml target\front-api.yaml

	$(GOBUILD) -o target/sys-rpc -v ./rpc/sys/sys.go
	$(GOBUILD) -o target/ums-rpc -v ./rpc/ums/ums.go
	$(GOBUILD) -o target/oms-rpc -v ./rpc/oms/oms.go
	$(GOBUILD) -o target/pms-rpc -v ./rpc/pms/pms.go
	$(GOBUILD) -o target/cms-rpc -v ./rpc/cms/cms.go
	$(GOBUILD) -o target/sms-rpc -v ./rpc/sms/sms.go
	$(GOBUILD) -o target/admin-api -v ./api/admin/admin.go
	$(GOBUILD) -o target/front-api -v ./api/front/front.go
	$(GOBUILD) -o target/web-api -v ./api/web/web.go


start: ## 运行目标
	nohup ./target/sys-rpc -f ./target/sys.yaml  > /dev/null 2>&1 &
	nohup ./target/ums-rpc -f ./target/ums.yaml  > /dev/null 2>&1 &
	nohup ./target/oms-rpc -f ./target/oms.yaml  > /dev/null 2>&1 &
	nohup ./target/pms-rpc -f ./target/pms.yaml  > /dev/null 2>&1 &
	nohup ./target/cms-rpc -f ./target/cms.yaml  > /dev/null 2>&1 &
	nohup ./target/sms-rpc -f ./target/sms.yaml  > /dev/null 2>&1 &
	nohup ./target/admin-api -f ./target/admin-api.yaml > /dev/null 2>&1 &
	nohup ./target/front-api -f ./target/front-api.yaml  > /dev/null 2>&1 &
	nohup ./target/web-api -f ./target/web-api.yaml  > /dev/null 2>&1 &


stop: ## 停止目标
	-pkill -f admin-api
	-pkill -f front-api
	-pkill -f web-api
	-pkill -f sys-rpc
	-pkill -f sms-rpc
	-pkill -f cms-rpc
	-pkill -f ums-rpc
	-pkill -f oms-rpc
	-pkill -f pms-rpc
	@for i in 5 4 3 2 1; do\
      echo -n "stop $$i";\
      sleep 1; \
      echo " "; \
    done



restart: stop start ## 重启项目


.DEFAULT_GOAL := all ## 默认构建目标是


GOCTL=$(GOBIN)/goctl ## goctl


format: ## 格式化代码
	$(GOCTL) api format --dir api/admin/doc/api
	$(GOCTL) api format --dir api/front/doc/api
	$(GOCTL) api format --dir api/web/doc/api

gen:	## 生成所有模块代码
	$(GOCTL) api go -api ./api/admin/doc/api/admin.api -dir ./api/admin/
	# 生成front-api代码
	$(GOCTL) api go -api ./api/front/doc/api/front.api -dir ./api/front/
	# 生成web-api代码
	$(GOCTL) api go -api ./api/web/doc/api/web.api -dir ./api/web/
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

model: ## 生成model代码
	$(GOCTL)  model mysql datasource -url="$(MYSQL_INFO)" -table="sys*" -dir=./rpc/model/sysmodel
	$(GOCTL)  model mysql datasource -url="$(MYSQL_INFO)" -table="ums*" -dir=./rpc/model/umsmodel
	$(GOCTL)  model mysql datasource -url="$(MYSQL_INFO)" -table="sms*" -dir=./rpc/model/smsmodel
	$(GOCTL)  model mysql datasource -url="$(MYSQL_INFO)" -table="oms*" -dir=./rpc/model/omsmodel
	$(GOCTL)  model mysql datasource -url="$(MYSQL_INFO)" -table="pms*" -dir=./rpc/model/pmsmodel
	$(GOCTL)  model mysql datasource -url="$(MYSQL_INFO)" -table="cms*" -dir=./rpc/model/cmsmodel

image: ## 构建docker镜像
	docker build -t sys-rpc:0.0.1 -f rpc/sys/Dockerfile .
	docker build -t ums-rpc:0.0.1 -f rpc/ums/Dockerfile .
	docker build -t oms-rpc:0.0.1 -f rpc/oms/Dockerfile .
	docker build -t pms-rpc:0.0.1 -f rpc/pms/Dockerfile .
	docker build -t sms-rpc:0.0.1 -f rpc/sms/Dockerfile .
	docker build -t cms-rpc:0.0.1 -f rpc/cms/Dockerfile .
	docker build -t admin-api:0.0.1 -f api/admin/Dockerfile .
	docker build -t front-api:0.0.1 -f api/front/Dockerfile .
	docker build -t web-api:0.0.1 -f api/web/Dockerfile .

run: ## 启动docker容器
	docker run -itd --net=host --name=sys sys-rpc:0.0.1; \
    docker run -itd --net=host --name=ums ums-rpc:0.0.1; \
    docker run -itd --net=host --name=oms oms-rpc:0.0.1; \
    docker run -itd --net=host --name=pms pms-rpc:0.0.1; \
    docker run -itd --net=host --name=sms sms-rpc:0.0.1; \
    docker run -itd --net=host --name=cms cms-rpc:0.0.1; \
    docker run -itd --net=host --name=admin-api admin-api:0.0.1; \
    docker run -itd --net=host --name=front-api front-api:0.0.1 \
    docker run -itd --net=host --name=web-api web-api:0.0.1 \

kubectl: ## 部署k8s容器
	kubectl apply -f script/account/serviceaccount.yaml; \
    kubectl apply -f script/configmap/register.yaml; \
    kubectl apply -f script/sys-rpc.yaml; \
    kubectl apply -f script/ums-rpc.yaml; \
    kubectl apply -f script/sms-rpc.yaml; \
    kubectl apply -f script/pms-rpc.yaml; \
    kubectl apply -f script/oms-rpc.yaml; \
    kubectl apply -f script/cms-rpc.yaml; \
    kubectl apply -f script/admin-api.yaml; \
    kubectl apply -f script/front-api.yaml; \
    kubectl apply -f script/web-api.yaml; \

help: ## show help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% 0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
