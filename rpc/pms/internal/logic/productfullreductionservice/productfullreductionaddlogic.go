package productfullreductionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductFullReductionAddLogic 产品满减
/*
Author: LiuFeiHua
Date: 2024/5/8 11:00
*/
type ProductFullReductionAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductFullReductionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductFullReductionAddLogic {
	return &ProductFullReductionAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductFullReductionAdd 添加产品满减
func (l *ProductFullReductionAddLogic) ProductFullReductionAdd(in *pmsclient.ProductFullReductionAddReq) (*pmsclient.ProductFullReductionAddResp, error) {
	err := query.PmsProductFullReduction.WithContext(l.ctx).Create(&model.PmsProductFullReduction{
		ProductID:   in.ProductId,
		FullPrice:   float64(in.FullPrice),
		ReducePrice: float64(in.ReducePrice),
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductFullReductionAddResp{}, nil
}
