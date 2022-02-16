package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuStockAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkuStockAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) SkuStockAddLogic {
	return SkuStockAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SkuStockAddLogic) SkuStockAdd(req types.AddSkuStockReq) (*types.AddSkuStockResp, error) {
	_, err := l.svcCtx.Pms.SkuStockAdd(l.ctx, &pmsclient.SkuStockAddReq{
		ProductId:      req.ProductId,
		SkuCode:        req.SkuCode,
		Price:          int64(req.Price),
		Stock:          req.Stock,
		LowStock:       req.LowStock,
		Pic:            req.Pic,
		Sale:           req.Sale,
		PromotionPrice: int64(req.PromotionPrice),
		LockStock:      req.LockStock,
		SpData:         req.SpData,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加商品库存信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加商品库存失败")
	}

	return &types.AddSkuStockResp{
		Code:    "000000",
		Message: "添加商品库存成功",
	}, nil
}
