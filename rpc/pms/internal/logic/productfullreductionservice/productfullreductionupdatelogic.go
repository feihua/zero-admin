package productfullreductionservicelogic

import (
	"context"
	"zero-admin/rpc/model/pmsmodel"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductFullReductionUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductFullReductionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductFullReductionUpdateLogic {
	return &ProductFullReductionUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductFullReductionUpdateLogic) ProductFullReductionUpdate(in *pmsclient.ProductFullReductionUpdateReq) (*pmsclient.ProductFullReductionUpdateResp, error) {
	err := l.svcCtx.PmsProductFullReductionModel.Update(l.ctx, &pmsmodel.PmsProductFullReduction{
		Id:          in.Id,
		ProductId:   in.ProductId,
		FullPrice:   float64(in.FullPrice),
		ReducePrice: float64(in.ReducePrice),
	})
	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductFullReductionUpdateResp{}, nil
}
