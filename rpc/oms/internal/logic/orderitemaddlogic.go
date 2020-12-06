package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderItemAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderItemAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderItemAddLogic {
	return &OrderItemAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderItemAddLogic) OrderItemAdd(in *oms.OrderItemAddReq) (*oms.OrderItemAddResp, error) {
	// todo: add your logic here and delete this line

	return &oms.OrderItemAddResp{}, nil
}
