package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

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
	// todo: add your logic here and delete this line
	err := l.svcCtx.UmsIntegrationChangeHistoryModel.Delete(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}
	return &umsclient.IntegrationChangeHistoryDeleteResp{}, nil
}
