package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderReturnReasonListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnReasonListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnReasonListLogic {
	return &OrderReturnReasonListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderReturnReasonListLogic) OrderReturnReasonList(in *oms.OrderReturnReasonListReq) (*oms.OrderReturnReasonListResp, error) {
	// todo: add your logic here and delete this line

	return &oms.OrderReturnReasonListResp{}, nil
}
