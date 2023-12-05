package main

import (
	"flag"
	"fmt"
	"zero-admin/rpc/ums/internal/config"
	growthchangehistoryserviceServer "zero-admin/rpc/ums/internal/server/growthchangehistoryservice"
	integrationchangehistoryserviceServer "zero-admin/rpc/ums/internal/server/integrationchangehistoryservice"
	integrationconsumesettingserviceServer "zero-admin/rpc/ums/internal/server/integrationconsumesettingservice"
	memberattentionservicer "zero-admin/rpc/ums/internal/server/memberattentionservice"
	memberlevelserviceServer "zero-admin/rpc/ums/internal/server/memberlevelservice"
	memberloginlogserviceServer "zero-admin/rpc/ums/internal/server/memberloginlogservice"
	membermembertagrelationserviceServer "zero-admin/rpc/ums/internal/server/membermembertagrelationservice"
	memberproductcategoryrelationserviceServer "zero-admin/rpc/ums/internal/server/memberproductcategoryrelationservice"
	memberproductcollectionserviceServer "zero-admin/rpc/ums/internal/server/memberproductcollectionservice"
	memberreadhistoryserviceServer "zero-admin/rpc/ums/internal/server/memberreadhistoryservice"
	memberreceiveaddressserviceServer "zero-admin/rpc/ums/internal/server/memberreceiveaddressservice"
	memberrulesettingserviceServer "zero-admin/rpc/ums/internal/server/memberrulesettingservice"
	memberserviceServer "zero-admin/rpc/ums/internal/server/memberservice"
	memberstatisticsinfoserviceServer "zero-admin/rpc/ums/internal/server/memberstatisticsinfoservice"
	membertaskserviceServer "zero-admin/rpc/ums/internal/server/membertaskservice"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "rpc/ums/etc/ums.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		umsclient.RegisterMemberServiceServer(grpcServer, memberserviceServer.NewMemberServiceServer(ctx))
		umsclient.RegisterGrowthChangeHistoryServiceServer(grpcServer, growthchangehistoryserviceServer.NewGrowthChangeHistoryServiceServer(ctx))
		umsclient.RegisterIntegrationChangeHistoryServiceServer(grpcServer, integrationchangehistoryserviceServer.NewIntegrationChangeHistoryServiceServer(ctx))
		umsclient.RegisterIntegrationConsumeSettingServiceServer(grpcServer, integrationconsumesettingserviceServer.NewIntegrationConsumeSettingServiceServer(ctx))
		umsclient.RegisterMemberLevelServiceServer(grpcServer, memberlevelserviceServer.NewMemberLevelServiceServer(ctx))
		umsclient.RegisterMemberLoginLogServiceServer(grpcServer, memberloginlogserviceServer.NewMemberLoginLogServiceServer(ctx))
		umsclient.RegisterMemberMemberTagRelationServiceServer(grpcServer, membermembertagrelationserviceServer.NewMemberMemberTagRelationServiceServer(ctx))
		umsclient.RegisterMemberProductCategoryRelationServiceServer(grpcServer, memberproductcategoryrelationserviceServer.NewMemberProductCategoryRelationServiceServer(ctx))
		umsclient.RegisterMemberReceiveAddressServiceServer(grpcServer, memberreceiveaddressserviceServer.NewMemberReceiveAddressServiceServer(ctx))
		umsclient.RegisterMemberRuleSettingServiceServer(grpcServer, memberrulesettingserviceServer.NewMemberRuleSettingServiceServer(ctx))
		umsclient.RegisterMemberStatisticsInfoServiceServer(grpcServer, memberstatisticsinfoserviceServer.NewMemberStatisticsInfoServiceServer(ctx))
		umsclient.RegisterMemberTaskServiceServer(grpcServer, membertaskserviceServer.NewMemberTaskServiceServer(ctx))
		umsclient.RegisterMemberProductCollectionServiceServer(grpcServer, memberproductcollectionserviceServer.NewMemberProductCollectionServiceServer(ctx))
		umsclient.RegisterMemberReadHistoryServiceServer(grpcServer, memberreadhistoryserviceServer.NewMemberReadHistoryServiceServer(ctx))
		umsclient.RegisterMemberAttentionServiceServer(grpcServer, memberattentionservicer.NewMemberAttentionServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
