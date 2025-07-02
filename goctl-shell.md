# 项目常用命令

## 1.相关工具安装

```shell
go get -u github.com/zeromicro/go-zero/tools/goctl
```

<font face="宋体" color=red size=3>以下操作,在项目根目录操作</font>

## 2.生成api代码

```shell
2.1后端管理代码
goctl api go -api ./api/admin/doc/api/admin.api -dir ./api/admin/

2.2移动端app代码
goctl api go -api ./api/front/doc/api/front.api -dir ./api/front/ --style go_zero

2.2网页端web代码
goctl api go -api ./api/web/doc/api/web.api -dir ./api/web/

2.3mq consumer代码
goctl api go -api ./consumer/consumer.api -dir ./consumer --style go_zero

2.4定时任务job代码
goctl api go -api ./job/job.api -dir ./job --style go_zero
```

## 3.生成rpc代码

```shell
3.1生成sys-rpc代码
goctl rpc protoc rpc/sys/sys.proto --go_out=./rpc/sys/ --go-grpc_out=./rpc/sys/ --zrpc_out=./rpc/sys/ -m
3.2生成ums-rpc代码
goctl rpc protoc rpc/ums/ums.proto --go_out=./rpc/ums/ --go-grpc_out=./rpc/ums/ --zrpc_out=./rpc/ums/ -m
3.3pms-rpc代码
goctl rpc protoc rpc/pms/pms.proto --go_out=./rpc/pms/ --go-grpc_out=./rpc/pms/ --zrpc_out=./rpc/pms/ -m
3.4生成oms-rpc代码
goctl rpc protoc rpc/oms/oms.proto --go_out=./rpc/oms/ --go-grpc_out=./rpc/oms/ --zrpc_out=./rpc/oms/ -m --style go_zero
3.5生成sms-rpc代码
goctl rpc protoc rpc/sms/sms.proto --go_out=./rpc/sms/ --go-grpc_out=./rpc/sms/ --zrpc_out=./rpc/sms/ -m --style go_zero
3.6生成cmsrpc代码
goctl rpc protoc rpc/cms/cms.proto --go_out=./rpc/cms/ --go-grpc_out=./rpc/cms/ --zrpc_out=./rpc/cms/ -m
3.7生成search代码
goctl rpc protoc rpc/search/search.proto --go_out=./rpc/search/ --go-grpc_out=./rpc/search/ --zrpc_out=./rpc/search/ --style go_zero 
```

## 3.生成mongo代码
```shell
goctl model mongo --type Test --dir .\rpc\ums\gen\mongo --style go_zero
```