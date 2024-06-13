package skustockservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteSkuStockLogic 删除sku的库存
/*
Author: LiuFeiHua
Date: 2024/6/12 17:11
*/
type DeleteSkuStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSkuStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSkuStockLogic {
	return &DeleteSkuStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteSkuStock 删除sku的库存
func (l *DeleteSkuStockLogic) DeleteSkuStock(in *pmsclient.DeleteSkuStockReq) (*pmsclient.DeleteSkuStockResp, error) {
	q := query.PmsSkuStock
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.DeleteSkuStockResp{}, nil
}
