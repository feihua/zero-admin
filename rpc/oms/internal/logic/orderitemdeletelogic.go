package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderItemDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderItemDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderItemDeleteLogic {
	return &OrderItemDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderItemDeleteLogic) OrderItemDelete(in *oms.OrderItemDeleteReq) (*oms.OrderItemDeleteResp, error) {
	err := l.svcCtx.OmsOrderItemModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &oms.OrderItemDeleteResp{}, nil
}
