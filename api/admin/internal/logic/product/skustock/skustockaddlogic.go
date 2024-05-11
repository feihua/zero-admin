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
	_, err := l.svcCtx.SkuStockService.SkuStockAdd(l.ctx, &pmsclient.SkuStockAddReq{
		ProductId:      req.ProductId,
		SkuCode:        req.SkuCode,
		Price:          req.Price,
		Stock:          req.Stock,
		LowStock:       req.LowStock,
		Pic:            req.Pic,
		Sale:           req.Sale,
		PromotionPrice: req.PromotionPrice,
		LockStock:      req.LockStock,
		SpData:         req.SpData,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加商品库存信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加商品库存失败")
	}

	return &types.AddSkuStockResp{
		Code:    "000000",
		Message: "添加商品库存成功",
	}, nil
}
