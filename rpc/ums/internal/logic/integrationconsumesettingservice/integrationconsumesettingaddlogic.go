package integrationconsumesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/umsmodel"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *IntegrationConsumeSettingAddLogic) IntegrationConsumeSettingAdd(in *umsclient.IntegrationConsumeSettingAddReq) (*umsclient.IntegrationConsumeSettingAddResp, error) {
	_, err := l.svcCtx.UmsIntegrationConsumeSettingModel.Insert(l.ctx, &umsmodel.UmsIntegrationConsumeSetting{
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
