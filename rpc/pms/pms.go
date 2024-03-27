package main

import (
	"flag"
	"fmt"

	"github.com/feihua/zero-admin/rpc/pms/internal/config"
	albumpicserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/albumpicservice"
	albumserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/albumservice"
	brandserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/brandservice"
	commentreplayserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/commentreplayservice"
	commentserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/commentservice"
	feighttemplateserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/feighttemplateservice"
	memberpriceserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/memberpriceservice"
	productattributecategoryserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productattributecategoryservice"
	productattributeserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productattributeservice"
	productattributevalueserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productattributevalueservice"
	productcategoryattributerelationserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productcategoryattributerelationservice"
	productcategoryserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productcategoryservice"
	productfullreductionserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productfullreductionservice"
	productladderserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productladderservice"
	productoperatelogserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productoperatelogservice"
	productserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productservice"
	productvertifyrecordserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/productvertifyrecordservice"
	skustockserviceServer "github.com/feihua/zero-admin/rpc/pms/internal/server/skustockservice"
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
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pmsclient.RegisterProductServiceServer(grpcServer, productserviceServer.NewProductServiceServer(ctx))
		pmsclient.RegisterAlbumServiceServer(grpcServer, albumserviceServer.NewAlbumServiceServer(ctx))
		pmsclient.RegisterAlbumPicServiceServer(grpcServer, albumpicserviceServer.NewAlbumPicServiceServer(ctx))
		pmsclient.RegisterBrandServiceServer(grpcServer, brandserviceServer.NewBrandServiceServer(ctx))
		pmsclient.RegisterCommentServiceServer(grpcServer, commentserviceServer.NewCommentServiceServer(ctx))
		pmsclient.RegisterCommentReplayServiceServer(grpcServer, commentreplayserviceServer.NewCommentReplayServiceServer(ctx))
		pmsclient.RegisterFeightTemplateServiceServer(grpcServer, feighttemplateserviceServer.NewFeightTemplateServiceServer(ctx))
		pmsclient.RegisterMemberPriceServiceServer(grpcServer, memberpriceserviceServer.NewMemberPriceServiceServer(ctx))
		pmsclient.RegisterProductAttributeCategoryServiceServer(grpcServer, productattributecategoryserviceServer.NewProductAttributeCategoryServiceServer(ctx))
		pmsclient.RegisterProductAttributeServiceServer(grpcServer, productattributeserviceServer.NewProductAttributeServiceServer(ctx))
		pmsclient.RegisterProductAttributeValueServiceServer(grpcServer, productattributevalueserviceServer.NewProductAttributeValueServiceServer(ctx))
		pmsclient.RegisterProductCategoryAttributeRelationServiceServer(grpcServer, productcategoryattributerelationserviceServer.NewProductCategoryAttributeRelationServiceServer(ctx))
		pmsclient.RegisterProductCategoryServiceServer(grpcServer, productcategoryserviceServer.NewProductCategoryServiceServer(ctx))
		pmsclient.RegisterProductFullReductionServiceServer(grpcServer, productfullreductionserviceServer.NewProductFullReductionServiceServer(ctx))
		pmsclient.RegisterProductLadderServiceServer(grpcServer, productladderserviceServer.NewProductLadderServiceServer(ctx))
		pmsclient.RegisterProductOperateLogServiceServer(grpcServer, productoperatelogserviceServer.NewProductOperateLogServiceServer(ctx))
		pmsclient.RegisterProductVertifyRecordServiceServer(grpcServer, productvertifyrecordserviceServer.NewProductVertifyRecordServiceServer(ctx))
		pmsclient.RegisterSkuStockServiceServer(grpcServer, skustockserviceServer.NewSkuStockServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
