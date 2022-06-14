package logic

import (
	"context"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderItemFastAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderItemFastAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderItemFastAddLogic {
	return &OrderItemFastAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderItemFastAddLogic) OrderItemFastAdd(in *omsclient.OrderItemFastAddReq) (*omsclient.OrderItemFastAddResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.OrderItemFastAddResp{}, nil
}
