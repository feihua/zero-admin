package integrationchangehistoryservicelogic

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *IntegrationChangeHistoryDeleteLogic) IntegrationChangeHistoryDelete(in *umsclient.IntegrationChangeHistoryDeleteReq) (*umsclient.IntegrationChangeHistoryDeleteResp, error) {
	err := l.svcCtx.UmsIntegrationChangeHistoryModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &umsclient.IntegrationChangeHistoryDeleteResp{}, nil
}
