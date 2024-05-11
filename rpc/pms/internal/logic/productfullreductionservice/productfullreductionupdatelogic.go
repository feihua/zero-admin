package productfullreductionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductFullReductionUpdateLogic 产品满减
/*
Author: LiuFeiHua
Date: 2024/5/8 11:01
*/
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

// ProductFullReductionUpdate 更新产品满减
func (l *ProductFullReductionUpdateLogic) ProductFullReductionUpdate(in *pmsclient.ProductFullReductionUpdateReq) (*pmsclient.ProductFullReductionUpdateResp, error) {
	q := query.PmsProductFullReduction
	_, err := q.WithContext(l.ctx).Updates(&model.PmsProductFullReduction{
		ID:          in.Id,
		ProductID:   in.ProductId,
		FullPrice:   in.FullPrice,
		ReducePrice: in.ReducePrice,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductFullReductionUpdateResp{}, nil
}
