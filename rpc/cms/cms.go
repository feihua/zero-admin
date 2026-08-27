package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/config"
	helpcategoryserviceServer "github.com/feihua/zero-admin/rpc/cms/internal/server/helpcategoryservice"
	helpserviceServer "github.com/feihua/zero-admin/rpc/cms/internal/server/helpservice"
	memberreportserviceServer "github.com/feihua/zero-admin/rpc/cms/internal/server/memberreportservice"
	preferredareaproductrelationserviceServer "github.com/feihua/zero-admin/rpc/cms/internal/server/preferredareaproductrelationservice"
	preferredareaserviceServer "github.com/feihua/zero-admin/rpc/cms/internal/server/preferredareaservice"
	subjectcategoryserviceServer "github.com/feihua/zero-admin/rpc/cms/internal/server/subjectcategoryservice"
	subjectcommentserviceServer "github.com/feihua/zero-admin/rpc/cms/internal/server/subjectcommentservice"
	subjectproductrelationserviceServer "github.com/feihua/zero-admin/rpc/cms/internal/server/subjectproductrelationservice"
	subjectserviceServer "github.com/feihua/zero-admin/rpc/cms/internal/server/subjectservice"
	topiccategoryserviceServer "github.com/feihua/zero-admin/rpc/cms/internal/server/topiccategoryservice"
	topiccommentserviceServer "github.com/feihua/zero-admin/rpc/cms/internal/server/topiccommentservice"
	topicserviceServer "github.com/feihua/zero-admin/rpc/cms/internal/server/topicservice"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile *string

// 初始化配置文件路径
func init() {
	defaultPath := "rpc/cms/etc/cms.yaml"
	configPath := os.Getenv("config_path")
	if strings.TrimSpace(configPath) != "" {
		if !strings.HasSuffix(configPath, "/") {
			configPath = configPath + "/"
		}
		defaultPath = configPath + defaultPath
	}
	configFile = flag.String("f", defaultPath, "the config file")
}

// 启动函数
func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)                     // 设置日志配置
	logx.AddWriter(logx.NewWriter(os.Stdout)) // 添加控制台输出
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		cmsclient.RegisterHelpServiceServer(grpcServer, helpserviceServer.NewHelpServiceServer(ctx))
		cmsclient.RegisterHelpCategoryServiceServer(grpcServer, helpcategoryserviceServer.NewHelpCategoryServiceServer(ctx))
		cmsclient.RegisterMemberReportServiceServer(grpcServer, memberreportserviceServer.NewMemberReportServiceServer(ctx))
		cmsclient.RegisterPreferredAreaServiceServer(grpcServer, preferredareaserviceServer.NewPreferredAreaServiceServer(ctx))
		cmsclient.RegisterPreferredAreaProductRelationServiceServer(grpcServer, preferredareaproductrelationserviceServer.NewPreferredAreaProductRelationServiceServer(ctx))
		cmsclient.RegisterSubjectServiceServer(grpcServer, subjectserviceServer.NewSubjectServiceServer(ctx))
		cmsclient.RegisterSubjectCategoryServiceServer(grpcServer, subjectcategoryserviceServer.NewSubjectCategoryServiceServer(ctx))
		cmsclient.RegisterSubjectCommentServiceServer(grpcServer, subjectcommentserviceServer.NewSubjectCommentServiceServer(ctx))
		cmsclient.RegisterSubjectProductRelationServiceServer(grpcServer, subjectproductrelationserviceServer.NewSubjectProductRelationServiceServer(ctx))
		cmsclient.RegisterTopicServiceServer(grpcServer, topicserviceServer.NewTopicServiceServer(ctx))
		cmsclient.RegisterTopicCategoryServiceServer(grpcServer, topiccategoryserviceServer.NewTopicCategoryServiceServer(ctx))
		cmsclient.RegisterTopicCommentServiceServer(grpcServer, topiccommentserviceServer.NewTopicCommentServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
