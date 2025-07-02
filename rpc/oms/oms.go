package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"

	"github.com/feihua/zero-admin/rpc/oms/internal/config"
	cartitemserviceServer "github.com/feihua/zero-admin/rpc/oms/internal/server/cartitemservice"
	companyaddressserviceServer "github.com/feihua/zero-admin/rpc/oms/internal/server/companyaddressservice"
	orderdeliveryserviceServer "github.com/feihua/zero-admin/rpc/oms/internal/server/orderdeliveryservice"
	orderoperationlogserviceServer "github.com/feihua/zero-admin/rpc/oms/internal/server/orderoperationlogservice"
	orderpaymentserviceServer "github.com/feihua/zero-admin/rpc/oms/internal/server/orderpaymentservice"
	orderreturnreasonserviceServer "github.com/feihua/zero-admin/rpc/oms/internal/server/orderreturnreasonservice"
	orderreturnserviceServer "github.com/feihua/zero-admin/rpc/oms/internal/server/orderreturnservice"
	orderserviceServer "github.com/feihua/zero-admin/rpc/oms/internal/server/orderservice"
	ordersettingserviceServer "github.com/feihua/zero-admin/rpc/oms/internal/server/ordersettingservice"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

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
	logx.MustSetup(c.Log)                     // 设置日志配置
	logx.AddWriter(logx.NewWriter(os.Stdout)) // 添加控制台输出
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		omsclient.RegisterCartItemServiceServer(grpcServer, cartitemserviceServer.NewCartItemServiceServer(ctx))
		omsclient.RegisterCompanyAddressServiceServer(grpcServer, companyaddressserviceServer.NewCompanyAddressServiceServer(ctx))
		omsclient.RegisterOrderDeliveryServiceServer(grpcServer, orderdeliveryserviceServer.NewOrderDeliveryServiceServer(ctx))
		omsclient.RegisterOrderOperationLogServiceServer(grpcServer, orderoperationlogserviceServer.NewOrderOperationLogServiceServer(ctx))
		omsclient.RegisterOrderPaymentServiceServer(grpcServer, orderpaymentserviceServer.NewOrderPaymentServiceServer(ctx))
		omsclient.RegisterOrderReturnReasonServiceServer(grpcServer, orderreturnreasonserviceServer.NewOrderReturnReasonServiceServer(ctx))
		omsclient.RegisterOrderReturnServiceServer(grpcServer, orderreturnserviceServer.NewOrderReturnServiceServer(ctx))
		omsclient.RegisterOrderServiceServer(grpcServer, orderserviceServer.NewOrderServiceServer(ctx))
		omsclient.RegisterOrderSettingServiceServer(grpcServer, ordersettingserviceServer.NewOrderSettingServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
