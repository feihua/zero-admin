package logic

import (
	"context"
	"go-zero-admin/rpc/model/umsmodel"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *IntegrationConsumeSettingAddLogic) IntegrationConsumeSettingAdd(in *ums.IntegrationConsumeSettingAddReq) (*ums.IntegrationConsumeSettingAddResp, error) {
	_, err := l.svcCtx.UmsIntegrationConsumeSettingModel.Insert(umsmodel.UmsIntegrationConsumeSetting{
		DeductionPerAmount: in.DeductionPerAmount,
		MaxPercentPerOrder: in.MaxPercentPerOrder,
		UseUnit:            in.UseUnit,
		CouponStatus:       in.CouponStatus,
	})
	if err != nil {
		return nil, err
	}

	return &ums.IntegrationConsumeSettingAddResp{}, nil
}
