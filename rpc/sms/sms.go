package main

import (
	"flag"
	"fmt"

	"zero-admin/rpc/sms/internal/config"
	couponhistoryserviceServer "zero-admin/rpc/sms/internal/server/couponhistoryservice"
	couponproductcategoryrelationserviceServer "zero-admin/rpc/sms/internal/server/couponproductcategoryrelationservice"
	couponproductrelationserviceServer "zero-admin/rpc/sms/internal/server/couponproductrelationservice"
	couponserviceServer "zero-admin/rpc/sms/internal/server/couponservice"
	flashpromotionlogserviceServer "zero-admin/rpc/sms/internal/server/flashpromotionlogservice"
	flashpromotionproductrelationserviceServer "zero-admin/rpc/sms/internal/server/flashpromotionproductrelationservice"
	flashpromotionserviceServer "zero-admin/rpc/sms/internal/server/flashpromotionservice"
	flashpromotionsessionserviceServer "zero-admin/rpc/sms/internal/server/flashpromotionsessionservice"
	homeadvertiseserviceServer "zero-admin/rpc/sms/internal/server/homeadvertiseservice"
	homebrandserviceServer "zero-admin/rpc/sms/internal/server/homebrandservice"
	homenewproductserviceServer "zero-admin/rpc/sms/internal/server/homenewproductservice"
	homerecommendproductserviceServer "zero-admin/rpc/sms/internal/server/homerecommendproductservice"
	homerecommendsubjectserviceServer "zero-admin/rpc/sms/internal/server/homerecommendsubjectservice"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "rpc/sms/etc/sms.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		smsclient.RegisterCouponServiceServer(grpcServer, couponserviceServer.NewCouponServiceServer(ctx))
		smsclient.RegisterCouponHistoryServiceServer(grpcServer, couponhistoryserviceServer.NewCouponHistoryServiceServer(ctx))
		smsclient.RegisterCouponProductCategoryRelationServiceServer(grpcServer, couponproductcategoryrelationserviceServer.NewCouponProductCategoryRelationServiceServer(ctx))
		smsclient.RegisterCouponProductRelationServiceServer(grpcServer, couponproductrelationserviceServer.NewCouponProductRelationServiceServer(ctx))
		smsclient.RegisterFlashPromotionServiceServer(grpcServer, flashpromotionserviceServer.NewFlashPromotionServiceServer(ctx))
		smsclient.RegisterFlashPromotionLogServiceServer(grpcServer, flashpromotionlogserviceServer.NewFlashPromotionLogServiceServer(ctx))
		smsclient.RegisterFlashPromotionProductRelationServiceServer(grpcServer, flashpromotionproductrelationserviceServer.NewFlashPromotionProductRelationServiceServer(ctx))
		smsclient.RegisterFlashPromotionSessionServiceServer(grpcServer, flashpromotionsessionserviceServer.NewFlashPromotionSessionServiceServer(ctx))
		smsclient.RegisterHomeAdvertiseServiceServer(grpcServer, homeadvertiseserviceServer.NewHomeAdvertiseServiceServer(ctx))
		smsclient.RegisterHomeBrandServiceServer(grpcServer, homebrandserviceServer.NewHomeBrandServiceServer(ctx))
		smsclient.RegisterHomeNewProductServiceServer(grpcServer, homenewproductserviceServer.NewHomeNewProductServiceServer(ctx))
		smsclient.RegisterHomeRecommendProductServiceServer(grpcServer, homerecommendproductserviceServer.NewHomeRecommendProductServiceServer(ctx))
		smsclient.RegisterHomeRecommendSubjectServiceServer(grpcServer, homerecommendsubjectserviceServer.NewHomeRecommendSubjectServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
