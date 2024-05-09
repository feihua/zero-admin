package skustockservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// SkuStockAddLogic 库存
/*
Author: LiuFeiHua
Date: 2024/5/8 11:10
*/
type SkuStockAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSkuStockAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuStockAddLogic {
	return &SkuStockAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SkuStockAdd 添加库存
func (l *SkuStockAddLogic) SkuStockAdd(in *pmsclient.SkuStockAddReq) (*pmsclient.SkuStockAddResp, error) {
	err := query.PmsSkuStock.WithContext(l.ctx).Create(&model.PmsSkuStock{
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

	return &pmsclient.SkuStockAddResp{}, nil
}
