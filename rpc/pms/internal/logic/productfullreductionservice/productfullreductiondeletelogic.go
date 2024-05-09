package productfullreductionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductFullReductionDeleteLogic 产品满减
/*
Author: LiuFeiHua
Date: 2024/5/8 9:59
*/
type ProductFullReductionDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductFullReductionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductFullReductionDeleteLogic {
	return &ProductFullReductionDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductFullReductionDelete 删除产品满减
func (l *ProductFullReductionDeleteLogic) ProductFullReductionDelete(in *pmsclient.ProductFullReductionDeleteReq) (*pmsclient.ProductFullReductionDeleteResp, error) {
	q := query.PmsProductFullReduction
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductFullReductionDeleteResp{}, nil
}
