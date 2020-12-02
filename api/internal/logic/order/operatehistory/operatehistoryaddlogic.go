package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OperateHistoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperateHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) OperateHistoryAddLogic {
	return OperateHistoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateHistoryAddLogic) OperateHistoryAdd(req types.AddOperateHistoryReq) (*types.AddOperateHistoryResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddOperateHistoryResp{}, nil
}
