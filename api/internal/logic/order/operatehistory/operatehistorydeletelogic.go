package logic

import (
	"context"
	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"
	"go-zero-admin/rpc/oms/omsclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type OperateHistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperateHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) OperateHistoryDeleteLogic {
	return OperateHistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateHistoryDeleteLogic) OperateHistoryDelete(req types.DeleteOperateHistoryReq) (*types.DeleteOperateHistoryResp, error) {
	_, _ = l.svcCtx.Oms.OrderOperateHistoryDelete(l.ctx, &omsclient.OrderOperateHistoryDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteOperateHistoryResp{
		Code:    "000000",
		Message: "",
	}, nil
}
