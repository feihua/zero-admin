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

// SkuStockUpdate 批量更新sku库存信息
func (l *SkuStockUpdateLogic) SkuStockUpdate(in *pmsclient.SkuStockUpdateReq) (*pmsclient.SkuStockUpdateResp, error) {
	var skuIds []int64
	var skuStockList []*model.PmsSkuStock
	for _, item := range in.SkuStockList {
		skuIds = append(skuIds, item.Id)
		skuStockList = append(skuStockList, &model.PmsSkuStock{
			ID:             item.Id,
			ProductID:      item.ProductId,
			SkuCode:        item.SkuCode,
			Price:          item.Price,
			Stock:          item.Stock,
			LowStock:       item.LowStock,
			Pic:            item.Pic,
			Sale:           item.Sale,
			PromotionPrice: item.PromotionPrice,
			LockStock:      item.LockStock,
			SpData:         item.SpData,
		})
	}
	q := query.PmsSkuStock
	//1.先删除
	_, err := q.WithContext(l.ctx).Where(q.ID.In(skuIds...)).Delete()
	if err != nil {
		return nil, err
	}

	//2.后添加
	err = q.WithContext(l.ctx).CreateInBatches(skuStockList, len(skuStockList))

	if err != nil {
		return nil, err
	}

	return &pmsclient.SkuStockUpdateResp{}, nil
}
