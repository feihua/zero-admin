package skustockservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QuerySkuStockListLogic 查询sku的库存列表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:12
*/
type QuerySkuStockListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySkuStockListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySkuStockListLogic {
	return &QuerySkuStockListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySkuStockList 查询sku的库存列表
func (l *QuerySkuStockListLogic) QuerySkuStockList(in *pmsclient.QuerySkuStockListReq) (*pmsclient.QuerySkuStockListResp, error) {
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
		})
	}

	return &pmsclient.QuerySkuStockListResp{
		Total: 0,
		List:  list,
	}, nil

}
