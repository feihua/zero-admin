package productfullreductionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductFullReductionLogic 添加产品满减表(只针对同商品)
/*
Author: LiuFeiHua
Date: 2024/6/12 17:05
*/
type AddProductFullReductionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductFullReductionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductFullReductionLogic {
	return &AddProductFullReductionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductFullReduction 添加产品满减表(只针对同商品)
func (l *AddProductFullReductionLogic) AddProductFullReduction(in *pmsclient.AddProductFullReductionReq) (*pmsclient.AddProductFullReductionResp, error) {
	err := query.PmsProductFullReduction.WithContext(l.ctx).Create(&model.PmsProductFullReduction{
		ProductID:   in.ProductId,   // 商品id
		FullPrice:   in.FullPrice,   // 商品满多少
		ReducePrice: in.ReducePrice, // 商品减多少
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.AddProductFullReductionResp{}, nil
}
