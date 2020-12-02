package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OperateHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperateHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) OperateHistoryListLogic {
	return OperateHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateHistoryListLogic) OperateHistoryList(req types.ListOperateHistoryReq) (*types.ListOperateHistoryResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListOperateHistoryResp{}, nil
}
