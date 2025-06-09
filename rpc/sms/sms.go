package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"

	"github.com/feihua/zero-admin/rpc/sms/internal/config"
	couponrecordserviceServer "github.com/feihua/zero-admin/rpc/sms/internal/server/couponrecordservice"
	couponscopeserviceServer "github.com/feihua/zero-admin/rpc/sms/internal/server/couponscopeservice"
	couponserviceServer "github.com/feihua/zero-admin/rpc/sms/internal/server/couponservice"
	coupontypeserviceServer "github.com/feihua/zero-admin/rpc/sms/internal/server/coupontypeservice"
	homeadvertiseserviceServer "github.com/feihua/zero-admin/rpc/sms/internal/server/homeadvertiseservice"
	seckillactivityserviceServer "github.com/feihua/zero-admin/rpc/sms/internal/server/seckillactivityservice"
	seckillproductserviceServer "github.com/feihua/zero-admin/rpc/sms/internal/server/seckillproductservice"
	seckillreservationserviceServer "github.com/feihua/zero-admin/rpc/sms/internal/server/seckillreservationservice"
	seckillsessionserviceServer "github.com/feihua/zero-admin/rpc/sms/internal/server/seckillsessionservice"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

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
	logx.MustSetup(c.Log)                     // 设置日志配置
	logx.AddWriter(logx.NewWriter(os.Stdout)) // 添加控制台输出
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		smsclient.RegisterCouponRecordServiceServer(grpcServer, couponrecordserviceServer.NewCouponRecordServiceServer(ctx))
		smsclient.RegisterCouponScopeServiceServer(grpcServer, couponscopeserviceServer.NewCouponScopeServiceServer(ctx))
		smsclient.RegisterCouponServiceServer(grpcServer, couponserviceServer.NewCouponServiceServer(ctx))
		smsclient.RegisterCouponTypeServiceServer(grpcServer, coupontypeserviceServer.NewCouponTypeServiceServer(ctx))
		smsclient.RegisterHomeAdvertiseServiceServer(grpcServer, homeadvertiseserviceServer.NewHomeAdvertiseServiceServer(ctx))
		smsclient.RegisterSeckillActivityServiceServer(grpcServer, seckillactivityserviceServer.NewSeckillActivityServiceServer(ctx))
		smsclient.RegisterSeckillProductServiceServer(grpcServer, seckillproductserviceServer.NewSeckillProductServiceServer(ctx))
		smsclient.RegisterSeckillReservationServiceServer(grpcServer, seckillreservationserviceServer.NewSeckillReservationServiceServer(ctx))
		smsclient.RegisterSeckillSessionServiceServer(grpcServer, seckillsessionserviceServer.NewSeckillSessionServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
