package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderReturnApplyUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnApplyUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnApplyUpdateLogic {
	return &OrderReturnApplyUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderReturnApplyUpdateLogic) OrderReturnApplyUpdate(in *oms.OrderReturnApplyUpdateReq) (*oms.OrderReturnApplyUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &oms.OrderReturnApplyUpdateResp{}, nil
}
