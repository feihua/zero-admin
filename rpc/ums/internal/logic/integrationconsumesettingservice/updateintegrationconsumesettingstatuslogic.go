package integrationconsumesettingservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateIntegrationConsumeSettingStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateIntegrationConsumeSettingStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateIntegrationConsumeSettingStatusLogic {
	return &UpdateIntegrationConsumeSettingStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新积分消费设置状态
func (l *UpdateIntegrationConsumeSettingStatusLogic) UpdateIntegrationConsumeSettingStatus(in *umsclient.UpdateIntegrationConsumeSettingStatusReq) (*umsclient.UpdateIntegrationConsumeSettingStatusResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.UpdateIntegrationConsumeSettingStatusResp{}, nil
}
