package integrationconsumesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// IntegrationConsumeSettingUpdateLogic 积分消费设置
/*
Author: LiuFeiHua
Date: 2024/5/7 10:03
*/
type IntegrationConsumeSettingUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationConsumeSettingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationConsumeSettingUpdateLogic {
	return &IntegrationConsumeSettingUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// IntegrationConsumeSettingUpdate 更新积分消费设置
func (l *IntegrationConsumeSettingUpdateLogic) IntegrationConsumeSettingUpdate(in *umsclient.IntegrationConsumeSettingUpdateReq) (*umsclient.IntegrationConsumeSettingUpdateResp, error) {
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

	return &umsclient.IntegrationConsumeSettingUpdateResp{}, nil
}
