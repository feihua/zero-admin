package orderreturnreasonservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderReturnReasonUpdateStatusLogic 退货原因
/*
Author: LiuFeiHua
Date: 2024/5/14 11:36
*/
type OrderReturnReasonUpdateStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnReasonUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnReasonUpdateStatusLogic {
	return &OrderReturnReasonUpdateStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderReturnReasonUpdateStatus 批量修改退货原因状态
func (l *OrderReturnReasonUpdateStatusLogic) OrderReturnReasonUpdateStatus(in *omsclient.OrderReturnReasonUpdateStatusReq) (*omsclient.OrderReturnReasonUpdateResp, error) {
	q := query.OmsOrderReturnReason
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		return nil, err
	}
	return &omsclient.OrderReturnReasonUpdateResp{}, nil
}
