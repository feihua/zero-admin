package logic

import (
	"context"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderItemCheckOutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderItemCheckOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderItemCheckOutLogic {
	return &OrderItemCheckOutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderItemCheckOutLogic) OrderItemCheckOut(in *omsclient.OrderItemCheckOutReq) (*omsclient.OrderItemCheckOutResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.OrderItemCheckOutResp{}, nil
}
