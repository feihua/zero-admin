package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type IntegrationConsumeSettingDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationConsumeSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationConsumeSettingDeleteLogic {
	return &IntegrationConsumeSettingDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IntegrationConsumeSettingDeleteLogic) IntegrationConsumeSettingDelete(in *umsclient.IntegrationConsumeSettingDeleteReq) (*umsclient.IntegrationConsumeSettingDeleteResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.UmsIntegrationConsumeSettingModel.Delete(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}
	return &umsclient.IntegrationConsumeSettingDeleteResp{}, nil
}
