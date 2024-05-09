package productladderservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductLadderUpdateLogic 产品阶梯价格
/*
Author: LiuFeiHua
Date: 2024/5/8 11:03
*/
type ProductLadderUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLadderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLadderUpdateLogic {
	return &ProductLadderUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductLadderUpdate 更新产品阶梯价格
func (l *ProductLadderUpdateLogic) ProductLadderUpdate(in *pmsclient.ProductLadderUpdateReq) (*pmsclient.ProductLadderUpdateResp, error) {
	q := query.PmsProductLadder
	_, err := q.WithContext(l.ctx).Updates(&model.PmsProductLadder{
		ID:        in.Id,
		ProductID: in.ProductId,
		Count:     in.Count,
		Discount:  float64(in.Discount),
		Price:     float64(in.Price),
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductLadderUpdateResp{}, nil
}
