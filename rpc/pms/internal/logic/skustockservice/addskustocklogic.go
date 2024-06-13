package skustockservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddSkuStockLogic 添加sku的库存
/*
Author: LiuFeiHua
Date: 2024/6/12 17:10
*/
type AddSkuStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSkuStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSkuStockLogic {
	return &AddSkuStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddSkuStock 添加sku的库存
func (l *AddSkuStockLogic) AddSkuStock(in *pmsclient.AddSkuStockReq) (*pmsclient.AddSkuStockResp, error) {
	err := query.PmsSkuStock.WithContext(l.ctx).Create(&model.PmsSkuStock{
		ProductID:      in.ProductId,
		SkuCode:        in.SkuCode,
		Price:          in.Price,
		Stock:          in.Stock,
		LowStock:       in.LowStock,
		Pic:            in.Pic,
		Sale:           in.Sale,
		PromotionPrice: in.PromotionPrice,
		LockStock:      in.LockStock,
		SpData:         in.SpData,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.AddSkuStockResp{}, nil
}
