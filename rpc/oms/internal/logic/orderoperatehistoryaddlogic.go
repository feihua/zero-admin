package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderOperateHistoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderOperateHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderOperateHistoryAddLogic {
	return &OrderOperateHistoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderOperateHistoryAddLogic) OrderOperateHistoryAdd(in *oms.OrderOperateHistoryAddReq) (*oms.OrderOperateHistoryAddResp, error) {
	// todo: add your logic here and delete this line

	return &oms.OrderOperateHistoryAddResp{}, nil
}
