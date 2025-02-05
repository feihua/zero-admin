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

// UpdateIntegrationConsumeSettingLogic 更新积分消费设置
/*
Author: LiuFeiHua
Date: 2024/5/23 9:38
*/
type UpdateIntegrationConsumeSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateIntegrationConsumeSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateIntegrationConsumeSettingLogic {
	return &UpdateIntegrationConsumeSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateIntegrationConsumeSetting 更新积分消费设置
func (l *UpdateIntegrationConsumeSettingLogic) UpdateIntegrationConsumeSetting(req *types.UpdateIntegrationConsumeSettingReq) (resp *types.UpdateIntegrationConsumeSettingResp, err error) {
	_, err = l.svcCtx.IntegrationConsumeSettingService.UpdateIntegrationConsumeSetting(l.ctx, &umsclient.UpdateIntegrationConsumeSettingReq{
		Id:                 req.Id,                 //
		DeductionPerAmount: req.DeductionPerAmount, // 每一元需要抵扣的积分数量
		MaxPercentPerOrder: req.MaxPercentPerOrder, // 每笔订单最高抵用百分比
		UseUnit:            req.UseUnit,            // 每次使用积分最小单位100
		IsDefault:          req.IsDefault,          // 是否默认：0->否；1->是
		CouponStatus:       req.CouponStatus,       // 是否可以和优惠券同用；0->不可以；1->可以
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新积分消费设置信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新积分消费设置失败")
	}

	return &types.UpdateIntegrationConsumeSettingResp{
		Code:    "000000",
		Message: "更新积分消费设置失败",
	}, nil
}
