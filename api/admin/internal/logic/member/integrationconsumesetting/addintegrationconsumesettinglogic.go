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

// AddIntegrationConsumeSettingLogic 添加积分消费设置
/*
Author: LiuFeiHua
Date: 2024/5/23 9:37
*/
type AddIntegrationConsumeSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddIntegrationConsumeSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIntegrationConsumeSettingLogic {
	return &AddIntegrationConsumeSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddIntegrationConsumeSetting 添加积分消费设置
func (l *AddIntegrationConsumeSettingLogic) AddIntegrationConsumeSetting(req *types.AddIntegrationConsumeSettingReq) (resp *types.AddIntegrationConsumeSettingResp, err error) {
	_, err = l.svcCtx.IntegrationConsumeSettingService.AddIntegrationConsumeSetting(l.ctx, &umsclient.AddIntegrationConsumeSettingReq{
		DeductionPerAmount: req.DeductionPerAmount,
		MaxPercentPerOrder: req.MaxPercentPerOrder,
		UseUnit:            req.UseUnit,
		CouponStatus:       req.CouponStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加积分消费设置信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加积分消费设置失败")
	}

	return &types.AddIntegrationConsumeSettingResp{
		Code:    "000000",
		Message: "添加积分消费设置成功",
	}, nil
}
