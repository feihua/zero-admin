package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderOperateHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderOperateHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderOperateHistoryListLogic {
	return &OrderOperateHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderOperateHistoryListLogic) OrderOperateHistoryList(in *oms.OrderOperateHistoryListReq) (*oms.OrderOperateHistoryListResp, error) {
	// todo: add your logic here and delete this line

	return &oms.OrderOperateHistoryListResp{}, nil
}
