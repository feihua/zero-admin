package integrationconsumesetting

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateIntegrationConsumeSettingStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateIntegrationConsumeSettingStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateIntegrationConsumeSettingStatusLogic {
	return &UpdateIntegrationConsumeSettingStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateIntegrationConsumeSettingStatusLogic) UpdateIntegrationConsumeSettingStatus(req *types.UpdateIntegrationConsumeSettingStatusReq) (resp *types.UpdateIntegrationConsumeSettingStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
