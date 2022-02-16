package logic

import (
	"context"
	"zero-admin/rpc/model/pmsmodel"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuStockUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSkuStockUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuStockUpdateLogic {
	return &SkuStockUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SkuStockUpdateLogic) SkuStockUpdate(in *pms.SkuStockUpdateReq) (*pms.SkuStockUpdateResp, error) {
	err := l.svcCtx.PmsSkuStockModel.Update(pmsmodel.PmsSkuStock{
		Id:             in.Id,
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

	return &pms.SkuStockUpdateResp{}, nil
}
