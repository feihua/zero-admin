package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/feihua/zero-admin/api/front/internal/config"
	"github.com/feihua/zero-admin/api/front/internal/handler"
	"github.com/feihua/zero-admin/api/front/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile *string

// 初始化配置文件路径
func init() {
	defaultPath := "api/front/etc/front-api.yaml"
	configPath := os.Getenv("config_path")
	if strings.TrimSpace(configPath) != "" {
		if !strings.HasSuffix(configPath, "/") {
			configPath = configPath + "/"
		}
		defaultPath = configPath + defaultPath
	}
	configFile = flag.String("f", defaultPath, "the config file")
}

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

	handler.RegisterHandlers(server, ctx)

	// httpx.SetErrorHandler(func(err error) (int, interface{}) {
	// 	var e *errorx.CodeError
	// 	switch {
	// 	case errors.As(err, &e):
	// 		return http.StatusOK, e.Data()
	// 	default:
	// 		return http.StatusOK, &errorx.CodeErrorResponse{
	// 			Code:    errorx.DefaultCode,
	// 			Message: e.Error(),
	// 		}
	// 	}
	// })

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	fmt.Printf("swagger ui at %s\n", "http://localhost:9999/swagger/index.html")
	server.Start()
}
