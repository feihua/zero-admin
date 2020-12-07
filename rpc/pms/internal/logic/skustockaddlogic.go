package logic

import (
	"context"
	"go-zero-admin/rpc/model/pmsmodel"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
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
