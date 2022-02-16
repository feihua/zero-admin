package logic

import (
	"context"
	"zero-admin/rpc/model/pmsmodel"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuStockAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSkuStockAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuStockAddLogic {
	return &SkuStockAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SkuStockAddLogic) SkuStockAdd(in *pms.SkuStockAddReq) (*pms.SkuStockAddResp, error) {
	_, err := l.svcCtx.PmsSkuStockModel.Insert(pmsmodel.PmsSkuStock{
		ProductId:      in.ProductId,
		SkuCode:        in.SkuCode,
		Price:          float64(in.Price),
		Stock:          in.Stock,
		LowStock:       in.LowStock,
		Pic:            in.Pic,
		Sale:           in.Sale,
		PromotionPrice: float64(in.PromotionPrice),
		LockStock:      in.LockStock,
		SpData:         in.SpData,
	})
	if err != nil {
		return nil, err
	}

	return &pms.SkuStockAddResp{}, nil
}
