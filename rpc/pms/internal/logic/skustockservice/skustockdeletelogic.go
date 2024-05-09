package skustockservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// SkuStockDeleteLogic 库存
/*
Author: LiuFeiHua
Date: 2024/5/8 10:05
*/
type SkuStockDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSkuStockDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuStockDeleteLogic {
	return &SkuStockDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SkuStockDelete 删除库存
func (l *SkuStockDeleteLogic) SkuStockDelete(in *pmsclient.SkuStockDeleteReq) (*pmsclient.SkuStockDeleteResp, error) {
	q := query.PmsSkuStock
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.SkuStockDeleteResp{}, nil
}
