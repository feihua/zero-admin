package logic

import (
	"context"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderConfirmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderConfirmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderConfirmLogic {
	return &OrderConfirmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderConfirmLogic) OrderConfirm(in *omsclient.OrderConfirmReq) (*omsclient.OrderConfirmResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.OrderConfirmResp{}, nil
}
