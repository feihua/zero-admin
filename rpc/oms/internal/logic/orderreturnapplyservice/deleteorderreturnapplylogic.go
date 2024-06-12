package orderreturnapplyservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteOrderReturnApplyLogic 删除订单退货申请
/*
Author: LiuFeiHua
Date: 2024/6/12 10:15
*/
type DeleteOrderReturnApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrderReturnApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderReturnApplyLogic {
	return &DeleteOrderReturnApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteOrderReturnApply 删除订单退货申请
func (l *DeleteOrderReturnApplyLogic) DeleteOrderReturnApply(in *omsclient.DeleteOrderReturnApplyReq) (*omsclient.DeleteOrderReturnApplyResp, error) {
	q := query.OmsOrderReturnApply
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &omsclient.DeleteOrderReturnApplyResp{}, nil
}
