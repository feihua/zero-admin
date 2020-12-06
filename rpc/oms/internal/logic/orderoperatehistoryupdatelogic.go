package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderOperateHistoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderOperateHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderOperateHistoryUpdateLogic {
	return &OrderOperateHistoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderOperateHistoryUpdateLogic) OrderOperateHistoryUpdate(in *oms.OrderOperateHistoryUpdateReq) (*oms.OrderOperateHistoryUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &oms.OrderOperateHistoryUpdateResp{}, nil
}
