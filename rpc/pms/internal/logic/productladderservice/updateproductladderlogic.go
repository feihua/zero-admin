package productladderservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateProductLadderLogic 更新产品阶梯价格表(只针对同商品)
/*
Author: LiuFeiHua
Date: 2024/6/12 17:08
*/
type UpdateProductLadderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLadderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLadderLogic {
	return &UpdateProductLadderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductLadder 更新产品阶梯价格表(只针对同商品)
func (l *UpdateProductLadderLogic) UpdateProductLadder(in *pmsclient.UpdateProductLadderReq) (*pmsclient.UpdateProductLadderResp, error) {
	q := query.PmsProductLadder
	_, err := q.WithContext(l.ctx).Updates(&model.PmsProductLadder{
		ID:        in.Id,
		ProductID: in.ProductId,
		Count:     in.Count,
		Discount:  in.Discount,
		Price:     in.Price,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.UpdateProductLadderResp{}, nil
}
