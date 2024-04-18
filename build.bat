copy rpc\sys\etc\sys.yaml target\sys.yaml
copy rpc\ums\etc\ums.yaml target\ums.yaml
copy rpc\oms\etc\oms.yaml target\oms.yaml
copy rpc\pms\etc\pms.yaml target\pms.yaml
copy rpc\cms\etc\cms.yaml target\cms.yaml
copy rpc\sms\etc\sms.yaml target\sms.yaml
copy api\admin\etc\admin-api.yaml target\admin-api.yaml
copy api\web\etc\web-api.yaml target\web-api.yaml
copy api\front\etc\front-api.yaml target\front-api.yaml


go build -o target/sys-rpc -v ./rpc/sys/sys.go
go build -o target/ums-rpc -v ./rpc/ums/ums.go
go build -o target/oms-rpc -v ./rpc/oms/oms.go
go build -o target/pms-rpc -v ./rpc/pms/pms.go
go build -o target/cms-rpc -v ./rpc/cms/cms.go
go build -o target/sms-rpc -v ./rpc/sms/sms.go
go build -o target/admin-api -v ./api/admin/admin.go
go build -o target/front-api -v ./api/front/front.go
go build -o target/web-api -v ./api/web/web.go

