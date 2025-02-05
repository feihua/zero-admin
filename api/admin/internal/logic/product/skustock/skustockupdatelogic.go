package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// SkuStockUpdateLogic 批量更新sku库存信息
/*
Author: LiuFeiHua
Date: 2024/5/15 15:46
*/
type SkuStockUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkuStockUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) SkuStockUpdateLogic {
	return SkuStockUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// SkuStockUpdate 批量更新sku库存信息
func (l *SkuStockUpdateLogic) SkuStockUpdate(req types.UpdateSkuStockReq) (*types.UpdateSkuStockResp, error) {
	var skuStockList []*pmsclient.UpdateSkuStockData
	for _, item := range req.SkuStockList {
		skuStockList = append(skuStockList, &pmsclient.UpdateSkuStockData{
			Id:             item.Id,             //
			ProductId:      item.ProductId,      // 商品id
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
	_, err := l.svcCtx.SkuStockService.UpdateSkuStock(l.ctx, &pmsclient.UpdateSkuStockReq{
		SkuStockList: skuStockList,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新商品库存信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新商品库存失败")
	}

	return &types.UpdateSkuStockResp{
		Code:    "000000",
		Message: "更新商品库存成功",
	}, nil
}
