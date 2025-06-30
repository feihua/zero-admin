package orderpaymentservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateOrderPaymentStatusLogic 更新订单支付
/*
Author: LiuFeiHua
Date: 2025/06/30 15:49:27
*/
type UpdateOrderPaymentStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderPaymentStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderPaymentStatusLogic {
	return &UpdateOrderPaymentStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrderPaymentStatus 更新订单支付状态
func (l *UpdateOrderPaymentStatusLogic) UpdateOrderPaymentStatus(in *omsclient.UpdateOrderPaymentStatusReq) (*omsclient.OrderPaymentResp, error) {
	q := query.OmsOrderPayment

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.PayStatus, in.PayStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新订单支付状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新订单支付状态失败")
	}

	return &omsclient.OrderPaymentResp{}, nil
}
