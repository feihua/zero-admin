package productladderservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductLadderLogic 添加产品阶梯价格表(只针对同商品)
/*
Author: LiuFeiHua
Date: 2024/6/12 17:07
*/
type AddProductLadderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductLadderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductLadderLogic {
	return &AddProductLadderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductLadder 添加产品阶梯价格表(只针对同商品)
func (l *AddProductLadderLogic) AddProductLadder(in *pmsclient.AddProductLadderReq) (*pmsclient.AddProductLadderResp, error) {
	err := query.PmsProductLadder.WithContext(l.ctx).Create(&model.PmsProductLadder{
		ProductID: in.ProductId,
		Count:     in.Count,
		Discount:  in.Discount,
		Price:     in.Price,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.AddProductLadderResp{}, nil
}
