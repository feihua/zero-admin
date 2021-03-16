package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GrowthChangeHistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrowthChangeHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) GrowthChangeHistoryDeleteLogic {
	return GrowthChangeHistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrowthChangeHistoryDeleteLogic) GrowthChangeHistoryDelete(req types.DeleteGrowthChangeHistoryReq) (*types.DeleteGrowthChangeHistoryResp, error) {
	_, _ = l.svcCtx.Ums.GrowthChangeHistoryDelete(l.ctx, &umsclient.GrowthChangeHistoryDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteGrowthChangeHistoryResp{
		Code:    "000000",
		Message: "",
	}, nil
}
