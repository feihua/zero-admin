package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/feihua/zero-admin/consumer/internal/config"
	"github.com/feihua/zero-admin/consumer/internal/handler"
	"github.com/feihua/zero-admin/consumer/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile *string

// 初始化配置文件路径
func init() {
	defaultPath := "consumer/etc/consumer-api.yaml"
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
