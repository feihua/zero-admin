1.安装
框架安装 go get github.com/tal-tech/go-zero
框架代码生成工具安装 go get -u github.com/tal-tech/go-zero/tools/goctl

2.创建api
进到api/doc/目录执行
goctl api -o borrow.api
goctl api go -api admin.api -dir ../

3.创建rpc
进到rpc/sys/目录操作
goctl rpc template -o sys.proto
goctl rpc proto -src sys.proto -dir .

进到rpc/ums/目录操作
goctl rpc template -o ums.proto
goctl rpc proto -src ums.proto -dir .

进到rpc/pms/目录操作
goctl rpc template -o pms.proto
goctl rpc proto -src pms.proto -dir .

进到rpc/oms/目录操作
goctl rpc template -o oms.proto
goctl rpc proto -src oms.proto -dir .

进到rpc/sms/目录操作
goctl rpc template -o sms.proto
goctl rpc proto -src sms.proto -dir .

4.创建model
进到rpc/目录操作
goctl model mysql ddl -c -src book.sql -dir .
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/gozero" -table="sys*" -dir ./model/sys

goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/gozero" -table="ums*" -dir ./model/ums
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/gozero" -table="pms*" -dir ./model/pms
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/gozero" -table="oms*" -dir ./model/oms
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/gozero" -table="sms*" -dir ./model/sms
