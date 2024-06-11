package integrationconsumesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateIntegrationConsumeSettingLogic 更新积分消费设置
/*
Author: LiuFeiHua
Date: 2024/6/11
*/
type UpdateIntegrationConsumeSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateIntegrationConsumeSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateIntegrationConsumeSettingLogic {
	return &UpdateIntegrationConsumeSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateIntegrationConsumeSetting 更新积分消费设置
func (l *UpdateIntegrationConsumeSettingLogic) UpdateIntegrationConsumeSetting(in *umsclient.UpdateIntegrationConsumeSettingReq) (*umsclient.UpdateIntegrationConsumeSettingResp, error) {
	_, err := query.UmsIntegrationConsumeSetting.WithContext(l.ctx).Updates(&model.UmsIntegrationConsumeSetting{
		ID:                 in.Id,
		DeductionPerAmount: in.DeductionPerAmount,
		MaxPercentPerOrder: in.MaxPercentPerOrder,
		UseUnit:            in.UseUnit,
		CouponStatus:       in.CouponStatus,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.UpdateIntegrationConsumeSettingResp{}, nil
}
