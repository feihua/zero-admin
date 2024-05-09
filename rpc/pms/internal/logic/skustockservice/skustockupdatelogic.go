package skustockservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// SkuStockUpdateLogic 库存
/*
Author: LiuFeiHua
Date: 2024/5/8 11:11
*/
type SkuStockUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSkuStockUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuStockUpdateLogic {
	return &SkuStockUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SkuStockUpdate 更新库存
func (l *SkuStockUpdateLogic) SkuStockUpdate(in *pmsclient.SkuStockUpdateReq) (*pmsclient.SkuStockUpdateResp, error) {
	q := query.PmsSkuStock
	_, err := q.WithContext(l.ctx).Updates(&model.PmsSkuStock{
		ID:             in.Id,
		ProductID:      in.ProductId,
		SkuCode:        in.SkuCode,
		Price:          float64(in.Price),
		Stock:          in.Stock,
		LowStock:       in.LowStock,
		Pic:            in.Pic,
		Sale:           in.Sale,
		PromotionPrice: float64(in.PromotionPrice),
		LockStock:      in.LockStock,
		SpData:         in.SpData,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.SkuStockUpdateResp{}, nil
}
