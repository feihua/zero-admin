package productfullreductionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductFullReductionLogic 删除产品满减表(只针对同商品)
/*
Author: LiuFeiHua
Date: 2024/6/12 17:05
*/
type DeleteProductFullReductionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductFullReductionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductFullReductionLogic {
	return &DeleteProductFullReductionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductFullReduction 删除产品满减表(只针对同商品)
func (l *DeleteProductFullReductionLogic) DeleteProductFullReduction(in *pmsclient.DeleteProductFullReductionReq) (*pmsclient.DeleteProductFullReductionResp, error) {
	q := query.PmsProductFullReduction
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.DeleteProductFullReductionResp{}, nil
}
