package main

import (
	"flag"
	"fmt"

	"zero-admin/rpc/oms/internal/config"
	cartitemserviceServer "zero-admin/rpc/oms/internal/server/cartitemservice"
	companyaddressserviceServer "zero-admin/rpc/oms/internal/server/companyaddressservice"
	orderitemserviceServer "zero-admin/rpc/oms/internal/server/orderitemservice"
	orderoperatehistorserviceServer "zero-admin/rpc/oms/internal/server/orderoperatehistorservice"
	orderreturnapplyserviceServer "zero-admin/rpc/oms/internal/server/orderreturnapplyservice"
	orderreturnreasonserviceServer "zero-admin/rpc/oms/internal/server/orderreturnreasonservice"
	orderserviceServer "zero-admin/rpc/oms/internal/server/orderservice"
	ordersettingserviceServer "zero-admin/rpc/oms/internal/server/ordersettingservice"
	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "rpc/oms/etc/oms.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		omsclient.RegisterOrderServiceServer(grpcServer, orderserviceServer.NewOrderServiceServer(ctx))
		omsclient.RegisterCartItemServiceServer(grpcServer, cartitemserviceServer.NewCartItemServiceServer(ctx))
		omsclient.RegisterCompanyAddressServiceServer(grpcServer, companyaddressserviceServer.NewCompanyAddressServiceServer(ctx))
		omsclient.RegisterOrderItemServiceServer(grpcServer, orderitemserviceServer.NewOrderItemServiceServer(ctx))
		omsclient.RegisterOrderOperateHistorServiceServer(grpcServer, orderoperatehistorserviceServer.NewOrderOperateHistorServiceServer(ctx))
		omsclient.RegisterOrderReturnApplyServiceServer(grpcServer, orderreturnapplyserviceServer.NewOrderReturnApplyServiceServer(ctx))
		omsclient.RegisterOrderReturnReasonServiceServer(grpcServer, orderreturnreasonserviceServer.NewOrderReturnReasonServiceServer(ctx))
		omsclient.RegisterOrderSettingServiceServer(grpcServer, ordersettingserviceServer.NewOrderSettingServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
