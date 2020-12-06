package logic

import (
	"context"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationChangeHistoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationChangeHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationChangeHistoryDeleteLogic {
	return &IntegrationChangeHistoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IntegrationChangeHistoryDeleteLogic) IntegrationChangeHistoryDelete(in *ums.IntegrationChangeHistoryDeleteReq) (*ums.IntegrationChangeHistoryDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &ums.IntegrationChangeHistoryDeleteResp{}, nil
}
