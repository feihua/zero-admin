package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"os"

	"github.com/feihua/zero-admin/job/internal/config"
	"github.com/feihua/zero-admin/job/internal/handler"
	"github.com/feihua/zero-admin/job/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "job/etc/job-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)                     // 设置日志配置
	logx.AddWriter(logx.NewWriter(os.Stdout)) // 添加控制台输出

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		var e *errorx.CodeError
		switch {
		case errors.As(err, &e):
			return http.StatusOK, e.Data()
		default:
			return http.StatusOK, &errorx.CodeErrorResponse{
				Code:    errorx.DefaultCode,
				Message: e.Error(),
			}
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
