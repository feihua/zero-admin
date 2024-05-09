package productladderservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductLadderDeleteLogic 产品阶梯价格
/*
Author: LiuFeiHua
Date: 2024/5/8 10:00
*/
type ProductLadderDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLadderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLadderDeleteLogic {
	return &ProductLadderDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductLadderDelete 删除产品阶梯价格
func (l *ProductLadderDeleteLogic) ProductLadderDelete(in *pmsclient.ProductLadderDeleteReq) (*pmsclient.ProductLadderDeleteResp, error) {
	q := query.PmsProductLadder
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductLadderDeleteResp{}, nil
}
