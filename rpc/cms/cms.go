package main

import (
	"flag"
	"fmt"

	"zero-admin/rpc/cms/cmsclient"
	"zero-admin/rpc/cms/internal/config"
	prefrenceareaproductrelationserviceServer "zero-admin/rpc/cms/internal/server/prefrenceareaproductrelationservice"
	prefrenceareaserviceServer "zero-admin/rpc/cms/internal/server/prefrenceareaservice"
	subjectproductrelationserviceServer "zero-admin/rpc/cms/internal/server/subjectproductrelationservice"
	subjectserviceServer "zero-admin/rpc/cms/internal/server/subjectservice"
	"zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "rpc/cms/etc/cms.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		cmsclient.RegisterSubjectServiceServer(grpcServer, subjectserviceServer.NewSubjectServiceServer(ctx))
		cmsclient.RegisterSubjectProductRelationServiceServer(grpcServer, subjectproductrelationserviceServer.NewSubjectProductRelationServiceServer(ctx))
		cmsclient.RegisterPrefrenceAreaServiceServer(grpcServer, prefrenceareaserviceServer.NewPrefrenceAreaServiceServer(ctx))
		cmsclient.RegisterPrefrenceAreaProductRelationServiceServer(grpcServer, prefrenceareaproductrelationserviceServer.NewPrefrenceAreaProductRelationServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
