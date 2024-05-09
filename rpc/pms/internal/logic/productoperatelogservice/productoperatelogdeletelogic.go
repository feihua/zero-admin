package productoperatelogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductOperateLogDeleteLogic 商品操作历史
/*
Author: LiuFeiHua
Date: 2024/5/8 10:01
*/
type ProductOperateLogDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductOperateLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductOperateLogDeleteLogic {
	return &ProductOperateLogDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductOperateLogDelete 删除商品操作历史
func (l *ProductOperateLogDeleteLogic) ProductOperateLogDelete(in *pmsclient.ProductOperateLogDeleteReq) (*pmsclient.ProductOperateLogDeleteResp, error) {
	q := query.PmsProductOperateLog
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductOperateLogDeleteResp{}, nil
}
