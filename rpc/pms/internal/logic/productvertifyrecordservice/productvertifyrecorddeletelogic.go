package productvertifyrecordservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductVertifyRecordDeleteLogic 商品审核
/*
Author: LiuFeiHua
Date: 2024/5/8 10:04
*/
type ProductVertifyRecordDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductVertifyRecordDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductVertifyRecordDeleteLogic {
	return &ProductVertifyRecordDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductVertifyRecordDelete 删除商品审核
func (l *ProductVertifyRecordDeleteLogic) ProductVertifyRecordDelete(in *pmsclient.ProductVertifyRecordDeleteReq) (*pmsclient.ProductVertifyRecordDeleteResp, error) {
	q := query.PmsProductVertifyRecord
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductVertifyRecordDeleteResp{}, nil
}
