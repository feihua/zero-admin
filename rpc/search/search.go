package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/feihua/zero-admin/rpc/search/internal/config"
	"github.com/feihua/zero-admin/rpc/search/internal/server"
	"github.com/feihua/zero-admin/rpc/search/internal/svc"
	"github.com/feihua/zero-admin/rpc/search/search"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile *string

// 初始化配置文件路径
func init() {
	defaultPath := "rpc/search/etc/search.yaml"
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

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		search.RegisterSearchServer(grpcServer, server.NewSearchServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
