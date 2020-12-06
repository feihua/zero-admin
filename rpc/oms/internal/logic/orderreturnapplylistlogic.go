package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderReturnApplyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnApplyListLogic {
	return &OrderReturnApplyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderReturnApplyListLogic) OrderReturnApplyList(in *oms.OrderReturnApplyListReq) (*oms.OrderReturnApplyListResp, error) {
	// todo: add your logic here and delete this line

	return &oms.OrderReturnApplyListResp{}, nil
}
