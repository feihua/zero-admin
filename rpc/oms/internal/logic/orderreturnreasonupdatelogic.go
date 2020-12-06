package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderReturnReasonUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnReasonUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnReasonUpdateLogic {
	return &OrderReturnReasonUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderReturnReasonUpdateLogic) OrderReturnReasonUpdate(in *oms.OrderReturnReasonUpdateReq) (*oms.OrderReturnReasonUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &oms.OrderReturnReasonUpdateResp{}, nil
}
