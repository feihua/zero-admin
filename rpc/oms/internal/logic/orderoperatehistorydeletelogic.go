package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderOperateHistoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderOperateHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderOperateHistoryDeleteLogic {
	return &OrderOperateHistoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderOperateHistoryDeleteLogic) OrderOperateHistoryDelete(in *oms.OrderOperateHistoryDeleteReq) (*oms.OrderOperateHistoryDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &oms.OrderOperateHistoryDeleteResp{}, nil
}
