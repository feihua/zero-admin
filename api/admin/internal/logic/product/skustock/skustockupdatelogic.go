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
	var skuStockList []*pmsclient.SkuStockUpdateData
	for _, item := range req.SkuStockList {
		skuStockList = append(skuStockList, &pmsclient.SkuStockUpdateData{
			Id:             item.Id,
			ProductId:      item.ProductId,
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
	_, err := l.svcCtx.SkuStockService.SkuStockUpdate(l.ctx, &pmsclient.SkuStockUpdateReq{
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
