package skustockservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// SkuStockListLogic 根据商品ID及sku编码模糊搜索sku库存
/*
Author: LiuFeiHua
Date: 2024/5/15 15:37
*/
type SkuStockListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSkuStockListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuStockListLogic {
	return &SkuStockListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SkuStockList 根据商品ID及sku编码模糊搜索sku库存
func (l *SkuStockListLogic) SkuStockList(in *pmsclient.SkuStockListReq) (*pmsclient.SkuStockListResp, error) {
	stock := query.PmsSkuStock
	q := stock.WithContext(l.ctx).Where(stock.ProductID.Eq(in.ProductId))
	if len(in.SkuCode) > 0 {
		q = q.Where(stock.SkuCode.Eq(in.SkuCode))
	}
	result, err := q.Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询库存列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.SkuStockListData
	for _, item := range result {

		list = append(list, &pmsclient.SkuStockListData{
			Id:             item.ID,
			ProductId:      item.ProductID,
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

	logc.Infof(l.ctx, "查询库存列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.SkuStockListResp{
		Total: 0,
		List:  list,
	}, nil
}
