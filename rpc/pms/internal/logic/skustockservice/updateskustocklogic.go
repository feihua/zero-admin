package skustockservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateSkuStockLogic 更新sku的库存
/*
Author: LiuFeiHua
Date: 2024/6/12 17:11
*/
type UpdateSkuStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSkuStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSkuStockLogic {
	return &UpdateSkuStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateSkuStock 更新sku的库存
func (l *UpdateSkuStockLogic) UpdateSkuStock(in *pmsclient.UpdateSkuStockReq) (*pmsclient.UpdateSkuStockResp, error) {
	var skuIds []int64
	var skuStockList []*model.PmsSkuStock
	for _, item := range in.SkuStockList {
		skuIds = append(skuIds, item.Id)
		skuStockList = append(skuStockList, &model.PmsSkuStock{
			ID:             item.Id,             //
			ProductID:      item.ProductId,      // 商品id
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
	q := query.PmsSkuStock
	// 1.先删除
	_, err := q.WithContext(l.ctx).Where(q.ID.In(skuIds...)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "更新sku的库存失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新sku的库存失败")
	}

	// 2.后添加
	err = q.WithContext(l.ctx).CreateInBatches(skuStockList, len(skuStockList))

	if err != nil {
		logc.Errorf(l.ctx, "更新sku的库存失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新sku的库存失败")
	}

	return &pmsclient.UpdateSkuStockResp{}, nil
}
