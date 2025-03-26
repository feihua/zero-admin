package integrationconsumesettingservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

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
		DeductionPerAmount: in.DeductionPerAmount, // 每一元需要抵扣的积分数量
		MaxPercentPerOrder: in.MaxPercentPerOrder, // 每笔订单最高抵用百分比
		UseUnit:            in.UseUnit,            // 每次使用积分最小单位100
		IsDefault:          in.IsDefault,          // 是否默认：0->否；1->是
		CouponStatus:       in.CouponStatus,       // 是否可以和优惠券同用；0->不可以；1->可以
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新积分消费设置失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新积分消费设置失败")
	}

	return &umsclient.UpdateIntegrationConsumeSettingResp{}, nil
}
