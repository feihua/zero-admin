package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"

	"github.com/feihua/zero-admin/rpc/pms/internal/config"
	commentreplayserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/commentreplayservice"
	commentserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/commentservice"
	feighttemplateserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/feighttemplateservice"
	productattributegroupserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productattributegroupservice"
	productattributeserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productattributeservice"
	productattributevalueserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productattributevalueservice"
	productbrandserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productbrandservice"
	productcategoryattributerelationserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productcategoryattributerelationservice"
	productcategoryserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productcategoryservice"
	productfullreductionserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productfullreductionservice"
	productladderserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productladderservice"
	productoperatelogserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productoperatelogservice"
	productskuserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productskuservice"
	productspecserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productspecservice"
	productspecvalueserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productspecvalueservice"
	productspuserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productspuservice"
	productvertifyrecordserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productvertifyrecordservice"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "rpc/pms/etc/pms.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)                     // 设置日志配置
	logx.AddWriter(logx.NewWriter(os.Stdout)) // 添加控制台输出
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pmsclient.RegisterProductSpuServiceServer(grpcServer, productspuserviceServer.NewProductSpuServiceServer(ctx))
		pmsclient.RegisterProductBrandServiceServer(grpcServer, productbrandserviceServer.NewProductBrandServiceServer(ctx))
		pmsclient.RegisterCommentServiceServer(grpcServer, commentserviceServer.NewCommentServiceServer(ctx))
		pmsclient.RegisterCommentReplayServiceServer(grpcServer, commentreplayserviceServer.NewCommentReplayServiceServer(ctx))
		pmsclient.RegisterFeightTemplateServiceServer(grpcServer, feighttemplateserviceServer.NewFeightTemplateServiceServer(ctx))
		pmsclient.RegisterProductAttributeGroupServiceServer(grpcServer, productattributegroupserviceServer.NewProductAttributeGroupServiceServer(ctx))
		pmsclient.RegisterProductAttributeServiceServer(grpcServer, productattributeserviceServer.NewProductAttributeServiceServer(ctx))
		pmsclient.RegisterProductAttributeValueServiceServer(grpcServer, productattributevalueserviceServer.NewProductAttributeValueServiceServer(ctx))
		pmsclient.RegisterProductCategoryAttributeRelationServiceServer(grpcServer, productcategoryattributerelationserviceServer.NewProductCategoryAttributeRelationServiceServer(ctx))
		pmsclient.RegisterProductCategoryServiceServer(grpcServer, productcategoryserviceServer.NewProductCategoryServiceServer(ctx))
		pmsclient.RegisterProductFullReductionServiceServer(grpcServer, productfullreductionserviceServer.NewProductFullReductionServiceServer(ctx))
		pmsclient.RegisterProductLadderServiceServer(grpcServer, productladderserviceServer.NewProductLadderServiceServer(ctx))
		pmsclient.RegisterProductOperateLogServiceServer(grpcServer, productoperatelogserviceServer.NewProductOperateLogServiceServer(ctx))
		pmsclient.RegisterProductVertifyRecordServiceServer(grpcServer, productvertifyrecordserviceServer.NewProductVertifyRecordServiceServer(ctx))
		pmsclient.RegisterProductSkuServiceServer(grpcServer, productskuserviceServer.NewProductSkuServiceServer(ctx))
		pmsclient.RegisterProductSpecServiceServer(grpcServer, productspecserviceServer.NewProductSpecServiceServer(ctx))
		pmsclient.RegisterProductSpecValueServiceServer(grpcServer, productspecvalueserviceServer.NewProductSpecValueServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
