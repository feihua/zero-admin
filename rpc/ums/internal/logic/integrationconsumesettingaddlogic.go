package logic

import (
	"context"

	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

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
	// todo: add your logic here and delete this line
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
