package integrationconsumesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// IntegrationConsumeSettingAddLogic 积分消费设置
/*
Author: LiuFeiHua
Date: 2024/5/7 10:01
*/
type IntegrationConsumeSettingAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationConsumeSettingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationConsumeSettingAddLogic {
	return &IntegrationConsumeSettingAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// IntegrationConsumeSettingAdd 添加积分消费设置
func (l *IntegrationConsumeSettingAddLogic) IntegrationConsumeSettingAdd(in *umsclient.IntegrationConsumeSettingAddReq) (*umsclient.IntegrationConsumeSettingAddResp, error) {
	err := query.UmsIntegrationConsumeSetting.WithContext(l.ctx).Create(&model.UmsIntegrationConsumeSetting{
		DeductionPerAmount: in.DeductionPerAmount,
		MaxPercentPerOrder: in.MaxPercentPerOrder,
		UseUnit:            in.UseUnit,
		CouponStatus:       in.CouponStatus,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.IntegrationConsumeSettingAddResp{}, nil
}
