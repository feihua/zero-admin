package skustockservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QuerySkuStockDetailLogic 查询sku的库存详情
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:07
*/
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

// QuerySkuStockDetail 查询sku的库存详情
func (l *QuerySkuStockDetailLogic) QuerySkuStockDetail(in *pmsclient.QuerySkuStockDetailReq) (*pmsclient.QuerySkuStockDetailResp, error) {
	item, err := query.PmsSkuStock.WithContext(l.ctx).Where(query.PmsSkuStock.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询sku的库存详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询sku的库存详情失败")
	}

	data := &pmsclient.QuerySkuStockDetailResp{
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
	}

	logc.Infof(l.ctx, "查询sku的库存详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
