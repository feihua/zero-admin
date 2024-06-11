package main

import (
	"flag"
	"fmt"
	"github.com/feihua/zero-admin/rpc/ums/internal/config"
	growthchangehistoryserviceServer "github.com/feihua/zero-admin/rpc/ums/internal/server/growthchangehistoryservice"
	integrationchangehistoryserviceServer "github.com/feihua/zero-admin/rpc/ums/internal/server/integrationchangehistoryservice"
	integrationconsumesettingserviceServer "github.com/feihua/zero-admin/rpc/ums/internal/server/integrationconsumesettingservice"
	memberbrandattentionserviceServer "github.com/feihua/zero-admin/rpc/ums/internal/server/memberbrandattentionservice"
	memberlevelserviceServer "github.com/feihua/zero-admin/rpc/ums/internal/server/memberlevelservice"
	memberloginlogserviceServer "github.com/feihua/zero-admin/rpc/ums/internal/server/memberloginlogservice"
	membermembertagrelationserviceServer "github.com/feihua/zero-admin/rpc/ums/internal/server/membermembertagrelationservice"
	memberproductcategoryrelationserviceServer "github.com/feihua/zero-admin/rpc/ums/internal/server/memberproductcategoryrelationservice"
	memberproductcollectionserviceServer "github.com/feihua/zero-admin/rpc/ums/internal/server/memberproductcollectionservice"
	memberreadhistoryserviceServer "github.com/feihua/zero-admin/rpc/ums/internal/server/memberreadhistoryservice"
	memberreceiveaddressserviceServer "github.com/feihua/zero-admin/rpc/ums/internal/server/memberreceiveaddressservice"
	memberrulesettingserviceServer "github.com/feihua/zero-admin/rpc/ums/internal/server/memberrulesettingservice"
	memberserviceServer "github.com/feihua/zero-admin/rpc/ums/internal/server/memberservice"
	memberstatisticsinfoserviceServer "github.com/feihua/zero-admin/rpc/ums/internal/server/memberstatisticsinfoservice"
	membertagserviceServer "github.com/feihua/zero-admin/rpc/ums/internal/server/membertagservice"
	membertaskserviceServer "github.com/feihua/zero-admin/rpc/ums/internal/server/membertaskservice"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

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
		umsclient.RegisterMemberTagServiceServer(grpcServer, membertagserviceServer.NewMemberTagServiceServer(ctx))
		umsclient.RegisterMemberProductCollectionServiceServer(grpcServer, memberproductcollectionserviceServer.NewMemberProductCollectionServiceServer(ctx))
		umsclient.RegisterMemberReadHistoryServiceServer(grpcServer, memberreadhistoryserviceServer.NewMemberReadHistoryServiceServer(ctx))
		umsclient.RegisterMemberBrandAttentionServiceServer(grpcServer, memberbrandattentionserviceServer.NewMemberBrandAttentionServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
