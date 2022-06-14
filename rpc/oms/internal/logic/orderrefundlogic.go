package logic

import (
	"context"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderRefundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderRefundLogic {
	return &OrderRefundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderRefundLogic) OrderRefund(in *omsclient.OrderRefundReq) (*omsclient.OrderRefundResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.OrderRefundResp{}, nil
}
