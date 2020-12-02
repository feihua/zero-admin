package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationChangeHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationChangeHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationChangeHistoryListLogic {
	return IntegrationChangeHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationChangeHistoryListLogic) IntegrationChangeHistoryList(req types.ListIntegrationChangeHistoryReq) (*types.ListIntegrationChangeHistoryResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListIntegrationChangeHistoryResp{}, nil
}
