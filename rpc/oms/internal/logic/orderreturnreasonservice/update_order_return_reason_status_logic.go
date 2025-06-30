package orderreturnreasonservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateOrderReturnReasonStatusLogic 更新退货原因
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
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

// UpdateOrderReturnReasonStatus 更新退货原因状态
func (l *UpdateOrderReturnReasonStatusLogic) UpdateOrderReturnReasonStatus(in *omsclient.UpdateOrderReturnReasonStatusReq) (*omsclient.UpdateOrderReturnReasonStatusResp, error) {
	q := query.OmsOrderReturnReason

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).UpdateSimple(q.Status.Value(in.Status), q.UpdateBy.Value(in.UpdateBy))

	if err != nil {
		logc.Errorf(l.ctx, "更新退货原因状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新退货原因状态失败")
	}

	return &omsclient.UpdateOrderReturnReasonStatusResp{}, nil
}
