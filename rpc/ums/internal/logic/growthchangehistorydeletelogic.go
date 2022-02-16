package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

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

func (l *GrowthChangeHistoryDeleteLogic) GrowthChangeHistoryDelete(in *ums.GrowthChangeHistoryDeleteReq) (*ums.GrowthChangeHistoryDeleteResp, error) {
	err := l.svcCtx.UmsGrowthChangeHistoryModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &ums.GrowthChangeHistoryDeleteResp{}, nil
}
