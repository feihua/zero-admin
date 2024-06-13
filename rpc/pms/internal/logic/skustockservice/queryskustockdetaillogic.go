package skustockservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySkuStockDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySkuStockDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySkuStockDetailLogic {
	return &QuerySkuStockDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询sku的库存详情
func (l *QuerySkuStockDetailLogic) QuerySkuStockDetail(in *pmsclient.QuerySkuStockDetailReq) (*pmsclient.QuerySkuStockDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pmsclient.QuerySkuStockDetailResp{}, nil
}
