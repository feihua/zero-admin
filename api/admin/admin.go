package main

import (
	"flag"
	"fmt"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/config"
	"github.com/feihua/zero-admin/api/admin/internal/handler"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"os"
)

var configFile = flag.String("f", "api/admin/etc/admin-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)                     // 设置日志配置
	logx.AddWriter(logx.NewWriter(os.Stdout)) // 添加控制台输出

	ctx := svc.NewServiceContext(c)
	var server *rest.Server
	if c.Swagger.IsTest {
		fs := rest.WithFileServer("/swagger", http.Dir(c.Swagger.Path))
		server = rest.MustNewServer(c.RestConf, fs)
	} else {
		server = rest.MustNewServer(c.RestConf)
	}

	defer server.Stop()

	server.Use(ctx.AddLog)

	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusOK, &errorx.CodeErrorResponse{
				Code:    errorx.DefaultCode,
				Message: e.Error(),
			}
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	fmt.Printf("swagger ui at %s\n", "http://localhost:8888/swagger/index.html")
	server.Start()
}
