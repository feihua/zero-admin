package skustockservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QuerySkuStockByProductSkuIdLogic
/*
Author: LiuFeiHua
Date: 2023/12/13 13:43
*/
type QuerySkuStockByProductSkuIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySkuStockByProductSkuIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySkuStockByProductSkuIdLogic {
	return &QuerySkuStockByProductSkuIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySkuStockByProductSkuId 根据ProductSkuId查询sku
func (l *QuerySkuStockByProductSkuIdLogic) QuerySkuStockByProductSkuId(in *pmsclient.QuerySkuStockByProductSkuIdReq) (*pmsclient.SkuStockListData, error) {
	q := query.PmsSkuStock
	stock, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.ProductSkuId)).First()
	if err != nil {
		return nil, err
	}

	return &pmsclient.SkuStockListData{
		Id:             stock.ID,
		ProductId:      stock.ProductID,
		SkuCode:        stock.SkuCode,
		Price:          float32(stock.Price),
		Stock:          stock.Stock,
		LowStock:       stock.LowStock,
		Pic:            stock.Pic,
		Sale:           stock.Sale,
		PromotionPrice: float32(stock.PromotionPrice),
		LockStock:      stock.LockStock,
		SpData:         stock.SpData,
	}, nil
}
