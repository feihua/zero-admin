package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderReturnApplyAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnApplyAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnApplyAddLogic {
	return &OrderReturnApplyAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderReturnApplyAddLogic) OrderReturnApplyAdd(in *oms.OrderReturnApplyAddReq) (*oms.OrderReturnApplyAddResp, error) {
	// todo: add your logic here and delete this line

	return &oms.OrderReturnApplyAddResp{}, nil
}
