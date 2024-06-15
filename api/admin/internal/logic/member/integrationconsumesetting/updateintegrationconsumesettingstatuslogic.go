package integrationconsumesetting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateIntegrationConsumeSettingStatusLogic 更新积分消费设置状态
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

// UpdateIntegrationConsumeSettingStatus 更新积分消费设置状态
func (l *UpdateIntegrationConsumeSettingStatusLogic) UpdateIntegrationConsumeSettingStatus(req *types.UpdateIntegrationConsumeSettingStatusReq) (resp *types.UpdateIntegrationConsumeSettingStatusResp, err error) {
	_, err = l.svcCtx.IntegrationConsumeSettingService.UpdateIntegrationConsumeSettingStatus(l.ctx, &umsclient.UpdateIntegrationConsumeSettingStatusReq{
		Id:        req.Id,
		IsDefault: req.IsDefault,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新积分消费设置状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新积分消费设置状态失败")
	}

	return &types.UpdateIntegrationConsumeSettingStatusResp{
		Code:    "000000",
		Message: "更新积分消费设置状态失败",
	}, nil

	return
}
