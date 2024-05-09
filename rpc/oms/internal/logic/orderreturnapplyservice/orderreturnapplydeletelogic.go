package orderreturnapplyservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderReturnApplyDeleteLogic 退货申请
/*
Author: LiuFeiHua
Date: 2024/5/8 9:33
*/
type OrderReturnApplyDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnApplyDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnApplyDeleteLogic {
	return &OrderReturnApplyDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderReturnApplyDelete 删除退货申请
func (l *OrderReturnApplyDeleteLogic) OrderReturnApplyDelete(in *omsclient.OrderReturnApplyDeleteReq) (*omsclient.OrderReturnApplyDeleteResp, error) {
	q := query.OmsOrderReturnApply
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &omsclient.OrderReturnApplyDeleteResp{}, nil
}
