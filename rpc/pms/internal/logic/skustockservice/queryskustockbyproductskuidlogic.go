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
	item, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.ProductSkuId)).First()
	if err != nil {
		return nil, err
	}

	return &pmsclient.SkuStockListData{
		Id:             item.ID,             //
		ProductId:      item.ProductID,      // 商品id
		SkuCode:        item.SkuCode,        // sku编码
		Price:          item.Price,          // 价格
		Stock:          item.Stock,          // 库存
		LowStock:       item.LowStock,       // 预警库存
		Pic:            item.Pic,            // 展示图片
		Sale:           item.Sale,           // 销量
		PromotionPrice: item.PromotionPrice, // 单品促销价格
		LockStock:      item.LockStock,      // 锁定库存
		SpData:         item.SpData,         // 商品销售属性，json格式
	}, nil
}
