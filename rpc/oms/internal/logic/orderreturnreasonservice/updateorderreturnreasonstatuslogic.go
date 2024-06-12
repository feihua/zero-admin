package orderreturnreasonservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateOrderReturnReasonStatusLogic 更新退货原因表状态
/*
Author: LiuFeiHua
Date: 2024/6/12 10:19
*/
type UpdateOrderReturnReasonStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderReturnReasonStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderReturnReasonStatusLogic {
	return &UpdateOrderReturnReasonStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrderReturnReasonStatus 更新退货原因表状态
func (l *UpdateOrderReturnReasonStatusLogic) UpdateOrderReturnReasonStatus(in *omsclient.UpdateOrderReturnReasonStatusReq) (*omsclient.UpdateOrderReturnReasonStatusResp, error) {
	q := query.OmsOrderReturnReason
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		return nil, err
	}

	return &omsclient.UpdateOrderReturnReasonStatusResp{}, nil
}
