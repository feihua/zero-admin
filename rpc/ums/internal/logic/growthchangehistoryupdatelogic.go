package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrowthChangeHistoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGrowthChangeHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrowthChangeHistoryUpdateLogic {
	return &GrowthChangeHistoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GrowthChangeHistoryUpdateLogic) GrowthChangeHistoryUpdate(in *umsclient.GrowthChangeHistoryUpdateReq) (*umsclient.GrowthChangeHistoryUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.GrowthChangeHistoryUpdateResp{}, nil
}
