package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// SkuStockListLogic 库存信息
/*
Author: LiuFeiHua
Date: 2024/5/15 15:33
*/
type SkuStockListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkuStockListLogic(ctx context.Context, svcCtx *svc.ServiceContext) SkuStockListLogic {
	return SkuStockListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// SkuStockList 根据商品ID及sku编码模糊搜索sku库存
func (l *SkuStockListLogic) SkuStockList(req types.ListSkuStockReq) (*types.ListSkuStockResp, error) {
	resp, err := l.svcCtx.SkuStockService.QuerySkuStockList(l.ctx, &pmsclient.QuerySkuStockListReq{
		ProductId: req.ProductId,
		SkuCode:   req.SkuCode,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询商品库存列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询商品库存失败")
	}

	var list []*types.ListSkuStockData

	for _, item := range resp.List {
		list = append(list, &types.ListSkuStockData{
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

	return &types.ListSkuStockResp{
		Data:    list,
		Code:    "000000",
		Message: "查询商品库存成功",
	}, nil
}
