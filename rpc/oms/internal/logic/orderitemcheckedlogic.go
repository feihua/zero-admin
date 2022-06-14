package logic

import (
	"context"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderItemCheckedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderItemCheckedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderItemCheckedLogic {
	return &OrderItemCheckedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderItemCheckedLogic) OrderItemChecked(in *omsclient.OrderItemCheckedReq) (*omsclient.OrderItemCheckedResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.OrderItemCheckedResp{}, nil
}
