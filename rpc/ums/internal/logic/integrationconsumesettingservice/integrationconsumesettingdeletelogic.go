package integrationconsumesettingservicelogic

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/rpc/ums/internal/svc"

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
	err := l.svcCtx.UmsIntegrationConsumeSettingModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &umsclient.IntegrationConsumeSettingDeleteResp{}, nil
}
