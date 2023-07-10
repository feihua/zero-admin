package growthchangehistoryservicelogic

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrowthChangeHistoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGrowthChangeHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrowthChangeHistoryDeleteLogic {
	return &GrowthChangeHistoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GrowthChangeHistoryDeleteLogic) GrowthChangeHistoryDelete(in *umsclient.GrowthChangeHistoryDeleteReq) (*umsclient.GrowthChangeHistoryDeleteResp, error) {
	err := l.svcCtx.UmsGrowthChangeHistoryModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &umsclient.GrowthChangeHistoryDeleteResp{}, nil
}
