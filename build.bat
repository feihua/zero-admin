IF "%1"=="sys-rpc" (
  go build -ldflags="-s -w"  -o target/sys-rpc/sys-rpc -v ./rpc/sys/sys.go
  copy rpc\sys\etc\sys.yaml target\sys-rpc\sys-rpc.yaml
  copy script\shell\start.sh target\sys-rpc\start.sh
) ELSE  IF "%1"=="ums-rpc" (
   go build -ldflags="-s -w" -o target/ums-rpc/ums-rpc -v ./rpc/ums/ums.go
   copy rpc\ums\etc\ums.yaml target\ums-rpc\ums-rpc.yaml
   copy script\shell\start.sh target\ums-rpc\start.sh
) ELSE  IF "%1"=="oms-rpc" (
   go build -ldflags="-s -w" -o target/oms-rpc/oms-rpc -v ./rpc/oms/oms.go
   copy rpc\oms\etc\oms.yaml target\oms-rpc\oms-rpc.yaml
   copy script\shell\start.sh target\oms-rpc\start.sh
) ELSE  IF "%1"=="pms-rpc" (
   go build -ldflags="-s -w" -o target/pms-rpc/pms-rpc -v ./rpc/pms/pms.go
   copy rpc\pms\etc\pms.yaml target\pms-rpc\pms-rpc.yaml
   copy script\shell\start.sh target\pms-rpc\start.sh
) ELSE  IF "%1"=="sms-rpc" (
   go build -ldflags="-s -w" -o target/sms-rpc/sms-rpc -v ./rpc/sms/sms.go
   copy rpc\sms\etc\sms.yaml target\sms-rpc\sms-rpc.yaml
   copy script\shell\start.sh target\sms-rpc\start.sh
) ELSE  IF "%1"=="cms-rpc" (
   go build -ldflags="-s -w" -o target/cms-rpc/cms-rpc -v ./rpc/cms/cms.go
   copy rpc\cms\etc\cms.yaml target\cms-rpc\cms-rpc.yaml
   copy script\shell\start.sh target\cms-rpc\start.sh
) ELSE  IF "%1"=="admin-api" (
   go build -ldflags="-s -w" -o target/admin-api/admin-api -v ./api/admin/admin.go
   copy api\admin\etc\admin-api.yaml target\admin-api\admin-api.yaml
   copy script\shell\start.sh target\admin-api\start.sh
) ELSE  IF "%1"=="web-api" (
   go build -ldflags="-s -w" -o target/web-api/web-api -v ./api/web/web.go
   copy api\web\etc\web-api.yaml target\web-api\web-api.yaml
   copy script\shell\start.sh target\web-api\start.sh
) ELSE  IF "%1"=="front-api" (
   go build -ldflags="-s -w" -o target/front-api/front-api -v ./api/front/front.go
   copy api\front\etc\front-api.yaml target\front-api\front-api.yaml
   copy script\shell\start.sh target\front-api\start.sh
) ELSE (
    go build -ldflags="-s -w" -o target/sys-rpc/sys-rpc -v ./rpc/sys/sys.go
    go build -ldflags="-s -w" -o target/ums-rpc/ums-rpc -v ./rpc/ums/ums.go
    go build -ldflags="-s -w" -o target/oms-rpc/oms-rpc -v ./rpc/oms/oms.go
    go build -ldflags="-s -w" -o target/pms-rpc/pms-rpc -v ./rpc/pms/pms.go
    go build -ldflags="-s -w" -o target/cms-rpc/cms-rpc -v ./rpc/cms/cms.go
    go build -ldflags="-s -w" -o target/sms-rpc/sms-rpc -v ./rpc/sms/sms.go
    go build -ldflags="-s -w" -o target/admin-api/admin-api -v ./api/admin/admin.go
    go build -ldflags="-s -w" -o target/front-api/front-api -v ./api/front/front.go
    go build -ldflags="-s -w" -o target/web-api/web-api -v ./api/web/web.go

    copy rpc\sys\etc\sys.yaml target\sys-rpc\sys-rpc.yaml
    copy rpc\ums\etc\ums.yaml target\ums-rpc\ums-rpc.yaml
    copy rpc\oms\etc\oms.yaml target\oms-rpc\oms-rpc.yaml
    copy rpc\pms\etc\pms.yaml target\pms-rpc\pms-rpc.yaml
    copy rpc\cms\etc\cms.yaml target\cms-rpc\cms-rpc.yaml
    copy rpc\sms\etc\sms.yaml target\sms-rpc\sms-rpc.yaml
    copy api\admin\etc\admin-api.yaml target\admin-api\admin-api.yaml
    copy api\web\etc\web-api.yaml target\web-api\web-api.yaml
    copy api\front\etc\front-api.yaml target\front-api\front-api.yaml

    copy script\shell\start.sh target\admin-api\start.sh
    copy script\shell\start.sh target\web-api\start.sh
    copy script\shell\start.sh target\front-api\start.sh
    copy script\shell\start.sh target\ums-rpc\start.sh
    copy script\shell\start.sh target\sys-rpc\start.sh
    copy script\shell\start.sh target\sms-rpc\start.sh
    copy script\shell\start.sh target\pms-rpc\start.sh
    copy script\shell\start.sh target\oms-rpc\start.sh
    copy script\shell\start.sh target\cms-rpc\start.sh
)





