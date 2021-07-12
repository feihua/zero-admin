package logic

import (
	"context"
	"encoding/json"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
		logx.Errorf("添加商品库存参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加商品库存失败")
	}

	return &types.AddSkuStockResp{
		Code:    "000000",
		Message: "添加商品库存成功",
	}, nil
}
