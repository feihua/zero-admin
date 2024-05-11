package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *SkuStockListLogic) SkuStockList(req types.ListSkuStockReq) (*types.ListSkuStockResp, error) {
	resp, err := l.svcCtx.SkuStockService.SkuStockList(l.ctx, &pmsclient.SkuStockListReq{
		ProductId: 1,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询商品库存列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询商品库存失败")
	}

	var list []*types.ListtSkuStockData

	for _, item := range resp.List {
		list = append(list, &types.ListtSkuStockData{
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

	return &types.ListSkuStockResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询商品库存成功",
	}, nil
}
