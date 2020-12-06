package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderItemUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderItemUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderItemUpdateLogic {
	return &OrderItemUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderItemUpdateLogic) OrderItemUpdate(in *oms.OrderItemUpdateReq) (*oms.OrderItemUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &oms.OrderItemUpdateResp{}, nil
}
